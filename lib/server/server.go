package server

import (
	"MyWebService/config"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type HandlerConfig struct {
	Path   string
	Method string
}

type Handler interface {
	HandleRequest(*gin.Context)
	GetConfig() HandlerConfig
}

func AsHandler(initHandlerFunc any) any {
	return fx.Annotate(
		initHandlerFunc,
		fx.As(new(Handler)),
		fx.ResultTags(`group:"handlers"`),
	)
}

type HTTPServerParams struct {
	fx.In

	Lc       fx.Lifecycle
	Cfg      *config.Config
	Handlers []Handler `group:"handlers"`
	Logger   *log.Logger
}

func NewHTTPServer(params HTTPServerParams) *http.Server {
	addr := fmt.Sprintf(":%d", params.Cfg.Port)
	router := gin.Default()
	srv := &http.Server{Addr: addr, Handler: router}

	for _, handler := range params.Handlers {
		router.Handle(handler.GetConfig().Method, handler.GetConfig().Path, handler.HandleRequest)
	}

	params.Lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			params.Logger.Printf("Starting %s server on %s\n", params.Cfg.Env, addr)

			go srv.ListenAndServe()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func StartHTTPServer(*http.Server) {
	// Does nothing, it is enough that this method is called via. FX with the http.Server argument.
	// Calling this method will trigger initialization of NewHTTPServer which has hook to start the server.
}
