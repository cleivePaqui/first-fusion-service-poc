package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/nable-fusion/fusion-api-template/internal/service"
	"github.com/nable-fusion/fusion-cloud-common/pkg/httputil"
	"net/http"
)

type Api struct {
	service service.ServiceInterface
}

func NewApi(service service.ServiceInterface) Api {
	return Api{
		service: service,
	}
}

// GetWelcomePage handles an API response (GET /).
func (i *Api) GetWelcomePage(ctx echo.Context) error {
	message := "Welcome to the Fusion Api Template!"
	err := ctx.String(http.StatusOK, message)
	httputil.AddContentLengthHeader(&ctx)
	return err
}

// GetHealthCheck handles an API response (GET /healthz).
func (i *Api) GetHealthCheck(ctx echo.Context) error {
	err := i.service.IntegrationPointsHealthCheck()
	if err != nil {
		return ctx.NoContent(http.StatusServiceUnavailable)
	}
	return ctx.NoContent(http.StatusNoContent)
}
