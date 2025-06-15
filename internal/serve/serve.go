package serve

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"jsonserver/internal/api/middlewares"
	apiValidator "jsonserver/internal/api/validator"
	"jsonserver/pkg/config"
)

type JsonServer struct {
	log *zap.SugaredLogger
}

func NewJsonServer(log *zap.SugaredLogger) *JsonServer {
	return &JsonServer{
		log: log,
	}
}

// ServeHTTP starts the HTTP server and returns the Echo instance.
func (s *JsonServer) ServeHTTP(port int64) *echo.Echo {
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

	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.log.Fatalf("shutting down the server. Err: %v", err)
	}

	return e
}
