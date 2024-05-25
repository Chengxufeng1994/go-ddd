package http

import (
	"fmt"
	"log"
	nethttp "net/http"
	"os"

	"github.com/Chengxufeng1994/go-ddd/config"
	"github.com/Chengxufeng1994/go-ddd/internal/application"
	"github.com/Chengxufeng1994/go-ddd/internal/transport/http/controller"
	"github.com/casbin/casbin/v2"
)

func NewHttpServer(cfg *config.Transport, enforcer *casbin.Enforcer, app *application.Application) *nethttp.Server {
	controller := controller.NewController(app)
	router := NewRouter(enforcer, controller)

	var handler nethttp.Handler = router

	addr := fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
	errStdLog := log.New(os.Stdout, "", log.LstdFlags)
	return &nethttp.Server{
		Addr:     addr,
		Handler:  handler,
		ErrorLog: errStdLog,
	}
}
