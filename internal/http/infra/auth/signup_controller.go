package auth

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/http/infra"
	auth_md "github.com/evgenivanovi/gomart/internal/model/auth"
	auth_uc "github.com/evgenivanovi/gomart/internal/usecase/auth"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/evgenivanovi/gomart/pkg/stdx/net/http/headers"
)

/* __________________________________________________ */

type SignUpRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

/* __________________________________________________ */

type SignUpController struct {
	decoder func(io.Reader) *json.Decoder
	usecase auth_uc.AuthSignUpUsecase
}

func ProvideSignUpController(
	usecase auth_uc.AuthSignUpUsecase,
) *SignUpController {
	decoder := func(reader io.Reader) *json.Decoder {
		decoder := json.NewDecoder(reader)
		decoder.DisallowUnknownFields()
		return decoder
	}
	return &SignUpController{
		decoder: decoder,
		usecase: usecase,
	}
}

func (c *SignUpController) Handle(
	writer http.ResponseWriter, request *http.Request,
) {

	requestModel, requestError := c.buildRequest(request)
	if requestError != nil {
		infra.LogErrorRequest(requestError)
		c.onError(requestError, nil, writer, request)
		return
	}

	responseModel, responseError := c.usecase.Execute(
		request.Context(), *requestModel,
	)

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

func (c *SignUpController) buildRequest(
	request *http.Request,
) (*auth_uc.SignUpRequest, error) {

	var requestModel SignUpRequest

	requestError := c.decoder(request.Body).Decode(&requestModel)
	if requestError != nil {
		return nil, requestError
	}

	return &auth_uc.SignUpRequest{
		Payload: auth_uc.SignUpRequestPayload{
			Credentials: auth_md.Credentials{
				Username: requestModel.Login,
				Password: requestModel.Password,
			},
		},
	}, nil

}

/* __________________________________________________ */

func (c *SignUpController) onSuccess(
	response auth_uc.SignUpResponse, writer http.ResponseWriter, request *http.Request,
) {
	writer.Header().Set(
		headers.AuthorizationKey.String(),
		headers.BuildBearerToken(response.Payload.Session.Tokens.AccessToken.Token),
	)
	writer.WriteHeader(http.StatusOK)
}

func (c *SignUpController) onError(
	requestError error, responseError error, writer http.ResponseWriter, request *http.Request,
) {
	if requestError != nil {
		c.translateRequestError(requestError, writer)
	}
	if responseError != nil {
		c.translateResponseError(responseError, writer)
	}
}

func (c *SignUpController) translateRequestError(
	err error, writer http.ResponseWriter,
) {
	writer.WriteHeader(http.StatusBadRequest)
}

func (c *SignUpController) translateResponseError(
	err error, writer http.ResponseWriter,
) {

	errCode := errx.ErrorCode(err)

	if errCode == core.ErrorExistsCode {
		writer.WriteHeader(http.StatusConflict)
		return
	}

	if errCode == errx.ErrorInternalCode {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

/* __________________________________________________ */
