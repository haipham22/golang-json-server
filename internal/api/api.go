package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"

	"jsonserver/internal/api/middlewares"
	apiValidator "jsonserver/internal/api/validator"
	"jsonserver/pkg/config"
)

type Handler struct {
	log *zap.SugaredLogger
}

func NewApiBiz(log *zap.SugaredLogger) *Handler {
	return &Handler{
		log: log,
	}
}

func (h *Handler) ServeHTTP() *echo.Echo {
	e := echo.New()

	e.Validator = apiValidator.NewCustomValidator()

	if config.ENV.APP.DEBUG {
		e.Debug = true
	}

	e.Use(
		middleware.RemoveTrailingSlashWithConfig(middleware.TrailingSlashConfig{
			RedirectCode: http.StatusMovedPermanently,
		}),
		middleware.Recover(),
		middleware.RequestID(),
		middlewares.Logger(zap.L()),
	)

	e.IPExtractor = echo.ExtractIPFromRealIPHeader()

	return e
}
