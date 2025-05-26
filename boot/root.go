package boot

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/kweusuf/url-shortner/configs"
	helloclient "github.com/kweusuf/url-shortner/pkg/client/hello"
	"github.com/kweusuf/url-shortner/pkg/constants"
	"github.com/kweusuf/url-shortner/pkg/endpoint"
	"github.com/kweusuf/url-shortner/pkg/service"
	helloservice "github.com/kweusuf/url-shortner/pkg/service/hello"
	"github.com/kweusuf/url-shortner/pkg/utils/log"
	httppkg "github.com/kweusuf/url-shortner/transport/http"
	"github.com/oklog/run"
)

type CancelInterrupt struct{}

func Init() {

	var conf *configs.AppConfig
	err := configs.InitializeConfig(conf)
	if err != nil {
		log.Error("Defaulting to inbuilt config")
		conf = configs.GetConfig()
	}

	ctx := initializeContext(nil) //TODO: Add support for config load here

	//TODO: Add logging logic here
	// Set global log level
	log.InitializeLogging()

	g := &run.Group{}

	endpoints := generateEndpoints(ctx)

	initializeHttpServer(endpoints, *conf, ctx, g)

	InitCancelInterrupt(g, make(chan CancelInterrupt))
	if err := g.Run(); err != nil {
		log.Error("Error in running go routines")
		log.Error(err.Error())
	}

}

func generateEndpoints(ctx context.Context) endpoint.AppEndpoints {
	helloClient := helloclient.MakeHelloClient()
	// mongoUtil := db.NewMongoUtil(ctx)
	// jobClient := jobclient.MakeJobManagerClient(*mongoUtil)

	helloService := helloservice.MakeHelloService(helloClient)
	// jobManagerService := jobservice.MakeJobManagerService(jobClient)

	endpoints := endpoint.MakeEndpoints(service.Services{
		HelloService: helloService,
		// JobManagerService: jobManagerService,
	})
	return endpoints
}

func initializeContext(config *configs.AppConfig) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.CtxKeyAppConfig{}, &config)
	return ctx
}

func initializeHttpServer(endpoints endpoint.AppEndpoints, config configs.AppConfig, ctx context.Context, g *run.Group) {
	httpHandler := httppkg.NewHttpHandler(endpoints)
	srv := http.Server{
		Addr:    config.URI.Port,
		Handler: httpHandler,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}
	g.Add(func() error {
		log.Info(fmt.Sprintf("Starting %s server at port %s", config.URI.HttpScheme, config.URI.Port))
		return srv.ListenAndServe()
	}, func(error) {
		log.Error(fmt.Sprintf("%s server exited", config.URI.HttpScheme))
		srv.Close()
	})
}

func InitCancelInterrupt(g *run.Group, cancelInterrupt chan CancelInterrupt) {
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}
