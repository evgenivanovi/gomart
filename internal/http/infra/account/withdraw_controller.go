package account

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/evgenivanovi/gomart/internal/domain/balance"
	"github.com/evgenivanovi/gomart/internal/http/infra"
	"github.com/evgenivanovi/gomart/internal/usecase/account"
	"github.com/evgenivanovi/gomart/internal/util/auth"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/pkg/errors"
)

/* __________________________________________________ */

type WithdrawRequest struct {
	Order string  `json:"order"`
	Sum   float64 `json:"sum"`
}

/* __________________________________________________ */

type WithdrawController struct {
	decoder func(io.Reader) *json.Decoder
	encoder func(io.Writer) *json.Encoder
	usecase account.CreateWithdrawUsecase
}

func ProvideWithdrawController(
	usecase account.CreateWithdrawUsecase,
) *WithdrawController {
	decoder := func(reader io.Reader) *json.Decoder {
		decoder := json.NewDecoder(reader)
		decoder.DisallowUnknownFields()
		return decoder
	}
	encoder := func(writer io.Writer) *json.Encoder {
		encoder := json.NewEncoder(writer)
		return encoder
	}
	return &WithdrawController{
		encoder: encoder,
		decoder: decoder,
		usecase: usecase,
	}
}

func (c *WithdrawController) Handle(
	writer http.ResponseWriter, request *http.Request,
) {

	requestModel, requestError := c.buildRequest(request)
	if requestError != nil {
		infra.LogErrorRequest(requestError)
		c.onError(requestError, nil, writer, request)
		return
	}

	responseError := c.usecase.Execute(request.Context(), *requestModel)

	if responseError != nil {
		infra.LogSuccessRequest(requestModel)
		infra.LogErrorResponse(responseError)
		c.onError(nil, responseError, writer, request)
	} else {
		infra.LogSuccessRequest(requestModel)
	}

}

func (c *WithdrawController) buildRequest(
	request *http.Request,
) (*account.WithdrawRequest, error) {

	User := auth.FromCtx(request.Context())
	if User == nil {
		return nil, errors.New("empty user")
	}

	var requestModel WithdrawRequest
	requestError := c.decoder(request.Body).Decode(&requestModel)
	if requestError != nil {
		return nil, requestError
	}

	return &account.WithdrawRequest{
		Payload: account.WithdrawRequestPayload{
			UserID: User.UserID,
			Order:  requestModel.Order,
			Amount: requestModel.Sum,
		},
	}, nil

}

/* __________________________________________________ */

func (c *WithdrawController) onSuccess(
	writer http.ResponseWriter, request *http.Request,
) {
	writer.WriteHeader(http.StatusOK)
}

func (c *WithdrawController) onError(
	requestError error, responseError error, writer http.ResponseWriter, request *http.Request,
) {
	if requestError != nil {
		c.translateRequestError(requestError, writer)
	}
	if responseError != nil {
		c.translateResponseError(responseError, writer)
	}
}

func (c *WithdrawController) translateRequestError(
	err error, writer http.ResponseWriter,
) {
	writer.WriteHeader(http.StatusBadRequest)
}

func (c *WithdrawController) translateResponseError(
	err error, writer http.ResponseWriter,
) {

	errCode := errx.ErrorCode(err)

	if errCode == balance.ErrorBalanceNotEnough {
		writer.WriteHeader(http.StatusPaymentRequired)
		return
	}

}

/* __________________________________________________ */
