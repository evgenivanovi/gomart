package account

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/evgenivanovi/gomart/internal/http/infra"
	"github.com/evgenivanovi/gomart/internal/usecase/account"
	"github.com/evgenivanovi/gomart/internal/util/auth"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/evgenivanovi/gomart/pkg/stdx/net/http/headers"
	"github.com/pkg/errors"
)

/* __________________________________________________ */

type GetBalanceResponse struct {
	Current   float64 `json:"current"`
	Withdrawn float64 `json:"withdrawn"`
}

/* __________________________________________________ */

type GetBalanceController struct {
	encoder func(io.Writer) *json.Encoder
	usecase account.GetBalanceUsecase
}

func ProvideGetBalanceController(
	usecase account.GetBalanceUsecase,
) *GetBalanceController {
	encoder := func(writer io.Writer) *json.Encoder {
		encoder := json.NewEncoder(writer)
		return encoder
	}
	return &GetBalanceController{
		encoder: encoder,
		usecase: usecase,
	}
}

func (c *GetBalanceController) Handle(
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

func (c *GetBalanceController) buildRequest(
	request *http.Request,
) (*account.GetBalanceRequest, error) {

	User := auth.FromCtx(request.Context())
	if User == nil {
		return nil, errors.New("empty user")
	}

	return &account.GetBalanceRequest{
		Payload: account.GetBalanceRequestPayload{
			UserID: User.UserID,
		},
	}, nil

}

/* __________________________________________________ */

func (c *GetBalanceController) onSuccess(
	response account.GetBalanceResponse, writer http.ResponseWriter, request *http.Request,
) {
	writer.Header().Set(
		headers.ContentTypeKey.String(),
		headers.TypeApplicationJSON.String(),
	)
	writer.WriteHeader(http.StatusOK)
	_ = c.encoder(writer).Encode(c.buildResponse(response))
}

func (c *GetBalanceController) buildResponse(
	response account.GetBalanceResponse,
) GetBalanceResponse {
	return GetBalanceResponse{
		Current:   response.Payload.Balance,
		Withdrawn: response.Payload.Withdrawn,
	}
}

func (c *GetBalanceController) onError(
	requestError error, responseError error, writer http.ResponseWriter, request *http.Request,
) {
	if requestError != nil {
		c.translateRequestError(requestError, writer)
	}
	if responseError != nil {
		c.translateResponseError(responseError, writer)
	}
}

func (c *GetBalanceController) translateRequestError(
	err error, writer http.ResponseWriter,
) {
	writer.WriteHeader(http.StatusBadRequest)
}

func (c *GetBalanceController) translateResponseError(
	err error, writer http.ResponseWriter,
) {

	errCode := errx.ErrorCode(err)

	if errCode == errx.ErrorInternalCode {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

/* __________________________________________________ */
