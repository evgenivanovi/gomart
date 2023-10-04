package order

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	order_dm "github.com/evgenivanovi/gomart/internal/domain/order"
	"github.com/evgenivanovi/gomart/internal/http/infra"
	"github.com/evgenivanovi/gomart/internal/usecase/order"
	"github.com/evgenivanovi/gomart/internal/util/auth"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	slices "github.com/evgenivanovi/gomart/pkg/std/slice"
	"github.com/evgenivanovi/gomart/pkg/stdx/net/http/headers"
	"github.com/pkg/errors"
)

/* __________________________________________________ */

type GetOrdersResponse struct {
	Number     string    `json:"number,omitempty"`
	Status     string    `json:"status,omitempty"`
	Accrual    float64   `json:"accrual,omitempty"`
	UploadedAt time.Time `json:"uploaded_at,omitempty"`
}

/* __________________________________________________ */

type GetOrdersController struct {
	encoder func(io.Writer) *json.Encoder
	usecase order.GetOrdersUsecase
}

func ProvideGetOrdersController(
	usecase order.GetOrdersUsecase,
) *GetOrdersController {
	encoder := func(writer io.Writer) *json.Encoder {
		encoder := json.NewEncoder(writer)
		return encoder
	}
	return &GetOrdersController{
		encoder: encoder,
		usecase: usecase,
	}
}

func (c *GetOrdersController) Handle(
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

func (c *GetOrdersController) buildRequest(
	request *http.Request,
) (*order.GetOrdersRequest, error) {

	User := auth.FromCtx(request.Context())
	if User == nil {
		return nil, errors.New("empty user")
	}

	return &order.GetOrdersRequest{
		Payload: order.GetOrdersRequestPayload{
			UserID: User.UserID,
		},
	}, nil

}

/* __________________________________________________ */

func (c *GetOrdersController) onSuccess(
	response order.GetOrdersResponse, writer http.ResponseWriter, request *http.Request,
) {
	if slices.IsEmpty(response.Payload.Orders) {
		writer.WriteHeader(http.StatusNoContent)
	} else {
		writer.Header().Set(
			headers.ContentTypeKey.String(),
			headers.TypeApplicationJSON.String(),
		)
		writer.WriteHeader(http.StatusOK)
		_ = c.encoder(writer).Encode(c.buildResponse(response))
	}
}

func (c *GetOrdersController) buildResponse(response order.GetOrdersResponse) []GetOrdersResponse {
	responses := make([]GetOrdersResponse, 0)
	for _, ord := range response.Payload.Orders {
		responses = append(
			responses,
			GetOrdersResponse{
				Number:     ord.ID,
				Status:     ord.Status,
				Accrual:    ord.Accrual,
				UploadedAt: ord.Metadata.CreatedAt,
			},
		)
	}
	return responses
}

func (c *GetOrdersController) onError(
	requestError error, responseError error, writer http.ResponseWriter, request *http.Request,
) {
	if requestError != nil {
		c.translateRequestError(requestError, writer)
	}
	if responseError != nil {
		c.translateResponseError(responseError, writer)
	}
}

func (c *GetOrdersController) translateRequestError(
	err error, writer http.ResponseWriter,
) {
	writer.WriteHeader(http.StatusBadRequest)
}

func (c *GetOrdersController) translateResponseError(
	err error, writer http.ResponseWriter,
) {

	errCode := errx.ErrorCode(err)

	if errCode == order_dm.ErrorOrderAlreadyLoaded {
		writer.WriteHeader(http.StatusOK)
		return
	}

	if errCode == order_dm.ErrorOrderAlreadyExists {
		writer.WriteHeader(http.StatusConflict)
		return
	}

	if errCode == errx.ErrorInternalCode {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

/* __________________________________________________ */
