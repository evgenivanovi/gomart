package order

import (
	"io"
	"net/http"

	order_dm "github.com/evgenivanovi/gomart/internal/domain/order"
	"github.com/evgenivanovi/gomart/internal/http/infra"
	"github.com/evgenivanovi/gomart/internal/usecase/order"
	"github.com/evgenivanovi/gomart/internal/util/auth"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/gookit/goutil/strutil"
	"github.com/pkg/errors"
)

/* __________________________________________________ */

type OrderLoadController struct {
	usecase order.OrderLoadUsecase
}

func ProvideOrderLoadController(
	usecase order.OrderLoadUsecase,
) *OrderLoadController {
	return &OrderLoadController{
		usecase: usecase,
	}
}

func (c *OrderLoadController) Handle(
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
		c.onSuccess(writer, request)
	}

}

func (c *OrderLoadController) buildRequest(
	request *http.Request,
) (*order.OrderLoadRequest, error) {

	requestBody, requestError := io.ReadAll(request.Body)
	if requestError != nil {
		return nil, requestError
	}

	OrderID := string(requestBody)
	if strutil.IsEmpty(OrderID) {
		return nil, errors.New("empty order")
	}

	User := auth.FromCtx(request.Context())
	if User == nil {
		return nil, errors.New("empty user")
	}

	return &order.OrderLoadRequest{
		Payload: order.OrderLoadRequestPayload{
			OrderID: OrderID,
			UserID:  User.UserID,
		},
	}, nil

}

/* __________________________________________________ */

func (c *OrderLoadController) onSuccess(
	writer http.ResponseWriter, request *http.Request,
) {
	writer.WriteHeader(http.StatusAccepted)
}

func (c *OrderLoadController) onError(
	requestError error, responseError error, writer http.ResponseWriter, request *http.Request,
) {
	if requestError != nil {
		c.translateRequestError(requestError, writer)
	}
	if responseError != nil {
		c.translateResponseError(responseError, writer)
	}
}

func (c *OrderLoadController) translateRequestError(
	err error, writer http.ResponseWriter,
) {
	writer.WriteHeader(http.StatusBadRequest)
}

func (c *OrderLoadController) translateResponseError(
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

	if errCode == order_dm.ErrorOrderInvalid {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if errCode == errx.ErrorInternalCode {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

}

/* __________________________________________________ */
