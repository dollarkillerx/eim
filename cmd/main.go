package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dollarkillerx/common/pkg/logger"
	"github.com/dollarkillerx/eim/internal/conf"
	"github.com/dollarkillerx/eim/internal/generated"
	"github.com/dollarkillerx/eim/internal/middlewares"
	"github.com/dollarkillerx/eim/internal/resolvers"
	"github.com/dollarkillerx/eim/internal/storage/simple"
	"github.com/dollarkillerx/eim/internal/utils"
	"github.com/go-chi/chi/v5"
)

var configFilename string
var configDirs string

func init() {
	const (
		defaultConfigFilename = "config"
		configUsage           = "Name of the config file, without extension"
		defaultConfigDirs     = "configs"
		configDirUsage        = "Directories to search for config file, separated by ','"
	)
	flag.StringVar(&configFilename, "c", defaultConfigFilename, configUsage)
	flag.StringVar(&configFilename, "config", defaultConfigFilename, configUsage)
	flag.StringVar(&configDirs, "cPath", defaultConfigDirs, configDirUsage)
}

func main() {
	eimInit()

	router := chi.NewRouter()
	router.Use(middlewares.Cors())
	if !conf.CONFIG.EnablePlayground {
		router.Use(middlewares.Safety())
	}

	router.Use(middlewares.Context())

	router.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	router.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ack"))
	})

	if conf.CONFIG.EnablePlayground {
		router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))

		log.Println("playground is enabled!")
		log.Printf("connect to http://%s/playground for GraphQL playground", conf.CONFIG.ListenAddress)
	}

	//将所有的rpcClient放入 gqlgen的Resolver
	newSimple, err := simple.NewSimple(conf.CONFIG.PostgresConfiguration)
	if err != nil {
		log.Fatalln(err)
	}
	cf := generated.Config{
		Resolvers: resolvers.NewResolver(newSimple),
	}

	cf.Directives.HasLogined = middlewares.HasLoginFunc

	graphQLServer := handler.NewDefaultServer(generated.NewExecutableSchema(cf))

	graphQLServer.SetRecoverFunc(middlewares.RecoverFunc)
	graphQLServer.SetErrorPresenter(middlewares.MiddleError)

	router.Handle("/graphql", graphQLServer)

	log.Printf("connect to http://%s/graphql", conf.CONFIG.ListenAddress)

	if err := http.ListenAndServe(conf.CONFIG.ListenAddress, router); err != nil {
		log.Fatalln(err)
	}
}

func eimInit() {
	utils.InitJWT()
	flag.Parse()

	// Setting up configurations
	err := conf.InitConfiguration(configFilename, configDirs)
	if err != nil {
		panic(fmt.Errorf("Error parsing config, %s", err))
	}

	// init logger
	logger.InitLogger(conf.CONFIG.LoggerConfig)
}
