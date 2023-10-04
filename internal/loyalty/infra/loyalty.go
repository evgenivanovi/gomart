package infra

import (
	"context"
	"fmt"
	"net/http"

	"github.com/evgenivanovi/gomart/internal/domain/core"
	"github.com/evgenivanovi/gomart/internal/domain/order"
	"github.com/evgenivanovi/gomart/internal/loyalty/adapter"
	errx "github.com/evgenivanovi/gomart/pkg/err"
	"github.com/go-resty/resty/v2"
)

/* __________________________________________________ */

type LoyaltyClientService struct {
	client *resty.Client
}

func ProvideLoyaltyClientService(client *resty.Client) *LoyaltyClientService {
	return &LoyaltyClientService{
		client: client,
	}
}

func (l *LoyaltyClientService) Get(
	ctx context.Context, request adapter.LoyaltyRequestModel,
) (*adapter.LoyaltyResponseModel, error) {
	url := fmt.Sprintf("/api/orders/%s", request.Order)

	resp, err := l.client.
		R().
		SetContext(ctx).
		SetResult(&adapter.LoyaltyResponseModel{}).
		Get(url)

	return l.getResponse(resp, err)
}

func (l *LoyaltyClientService) getResponse(
	response *resty.Response, err error,
) (*adapter.LoyaltyResponseModel, error) {

	if err != nil {
		return nil, errx.NewErrorWithEntityCodeMessage(
			order.LoyaltyEntityError, errx.ErrorInternalMessage, err.Error(),
		)
	}

	status := response.StatusCode()

	if status == http.StatusNoContent {
		return nil, errx.NewErrorWithEntityCode(
			order.LoyaltyEntityError, core.ErrorNotFoundCode,
		)
	}

	if status == http.StatusOK {
		if responseModel, ok := response.Result().(*adapter.LoyaltyResponseModel); ok {
			return responseModel, nil
		} else {
			return nil, errx.NewErrorWithEntityCode(
				order.LoyaltyEntityError, errx.ErrorInternalMessage,
			)
		}
	}

	return nil, errx.NewErrorWithEntityCode(
		order.LoyaltyEntityError, errx.ErrorInternalMessage,
	)

}

/* __________________________________________________ */
