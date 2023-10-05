package account

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/evgenivanovi/gomart/internal/domain/balance"
	"github.com/evgenivanovi/gomart/internal/http/infra"
	"github.com/evgenivanovi/gomart/internal/usecase/account"
	"github.com/evgenivanovi/gomart/internal/util/auth"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	slices "github.com/evgenivanovi/gomart/pkg/std/slice"
	"github.com/evgenivanovi/gomart/pkg/stdx/net/http/headers"
	"github.com/pkg/errors"
)

/* __________________________________________________ */

type WithdrawalHistoryResponse struct {
	Order       string    `json:"order"`
	Sum         float64   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

/* __________________________________________________ */

type WithdrawalHistoryController struct {
	decoder func(io.Reader) *json.Decoder
	encoder func(io.Writer) *json.Encoder
	usecase account.GetWithdrawalHistoryUsecase
}

func ProvideWithdrawalHistoryController(
	usecase account.GetWithdrawalHistoryUsecase,
) *WithdrawalHistoryController {
	decoder := func(reader io.Reader) *json.Decoder {
		decoder := json.NewDecoder(reader)
		decoder.DisallowUnknownFields()
		return decoder
	}
	encoder := func(writer io.Writer) *json.Encoder {
		encoder := json.NewEncoder(writer)
		return encoder
	}
	return &WithdrawalHistoryController{
		encoder: encoder,
		decoder: decoder,
		usecase: usecase,
	}
}

func (c *WithdrawalHistoryController) Handle(
	writer http.ResponseWriter, request *http.Request,
) {

	requestModel, requestError := c.buildRequest(request)
	if requestError != nil {
		infra.LogErrorRequest(requestError)
		c.onError(requestError, nil, writer, request)
		return
	}

	responseModel, responseError := c.usecase.Execute(request.Context(), *requestModel)

	if responseError != nil {
		infra.LogSuccessRequest(requestModel)
		infra.LogErrorResponse(responseError)
		c.onError(nil, responseError, writer, request)
	} else {
		infra.LogSuccessRequest(requestModel)
		infra.LogSuccessResponse(responseModel)
		c.onSuccess(responseModel, writer, request)
	}

}

func (c *WithdrawalHistoryController) buildRequest(
	request *http.Request,
) (*account.WithdrawalHistoryRequest, error) {

	User := auth.FromCtx(request.Context())
	if User == nil {
		return nil, errors.New("empty user")
	}

	return &account.WithdrawalHistoryRequest{
		Payload: account.WithdrawalHistoryRequestPayload{
			UserID: User.UserID,
		},
	}, nil

}

/* __________________________________________________ */

func (c *WithdrawalHistoryController) onSuccess(
	response account.WithdrawalHistoryResponse, writer http.ResponseWriter, request *http.Request,
) {

	writer.Header().Set(
		headers.ContentTypeKey.String(),
		headers.TypeApplicationJSON.String(),
	)

	if slices.IsEmpty(response.Payload.Withdrawals.Withdrawals) {
		writer.WriteHeader(http.StatusNoContent)
		return
	}

	writer.WriteHeader(http.StatusOK)
	_ = c.encoder(writer).Encode(c.buildResponse(response))

}

func (c *WithdrawalHistoryController) buildResponse(
	response account.WithdrawalHistoryResponse,
) []WithdrawalHistoryResponse {

	responses := make([]WithdrawalHistoryResponse, 0)

	for _, elem := range response.Payload.Withdrawals.Withdrawals {
		resp := WithdrawalHistoryResponse{
			Order:       elem.Order,
			Sum:         elem.Amount,
			ProcessedAt: elem.Metadata.CreatedAt,
		}
		responses = append(responses, resp)
	}

	return responses

}

func (c *WithdrawalHistoryController) onError(
	requestError error, responseError error, writer http.ResponseWriter, request *http.Request,
) {
	if requestError != nil {
		c.translateRequestError(requestError, writer)
	}
	if responseError != nil {
		c.translateResponseError(responseError, writer)
	}
}

func (c *WithdrawalHistoryController) translateRequestError(
	err error, writer http.ResponseWriter,
) {
	writer.WriteHeader(http.StatusBadRequest)
}

func (c *WithdrawalHistoryController) translateResponseError(
	err error, writer http.ResponseWriter,
) {

	errCode := errx.ErrorCode(err)

	if errCode == balance.ErrorBalanceNotEnough {
		writer.WriteHeader(http.StatusPaymentRequired)
		return
	}

}

/* __________________________________________________ */
