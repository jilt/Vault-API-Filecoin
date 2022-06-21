package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jilt/Vault-API-Filecoin/internal/handlers"
	"github.com/jilt/Vault-API-Filecoin/internal/utils"
	"github.com/jilt/Vault-API-Filecoin/logger"
	"github.com/spf13/viper"
)

func main() {
	// Create a deadline to wait for.
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*5, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := gin.New()

	cors := handlers.Cors{}
	r.Use(cors.Middleware())

	ownedHandler := handlers.OwnedHandler{}
	r.GET("owned/:user", ownedHandler.Handle)

	ownersHandler := handlers.OwnersHandler{}
	r.GET("owners/:tokenid", ownersHandler.Handle)

	ownedFilteredHandler := handlers.OwnedFilteredHandler{}
	r.GET("owned/:user/:store", ownedFilteredHandler.Handle)

	FmHandler := handlers.FmHandler{}
	r.GET("fm/:tokenid", fmHandler.Handle)
	
	ownedParasHandler := handlers.OwnedParasHandler{}
	r.GET("owned-paras/:user", ownedParasHandler.Handle)

	ownersParasHandler := handlers.OwnersParasHandler{}
	r.GET("owners-paras/:tokenid", ownersParasHandler.Handle)

	healthHandler := handlers.HealthHandler{}
	r.GET("healthz", healthHandler.Handle)

	UnlockableHandler := handlers.UnlockableHandler{}
	r.GET("unlockable/:tokenid", UnlockableHandler.Handle)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {

		if err := srv.ListenAndServe(); err != nil {
			log.Printf("server listen error, %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.

	c := make(chan os.Signal, 1)

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	log.Printf("shutdown server")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	os.Exit(0)

}

func init() {
	viper.AutomaticEnv()

	viper.SetEnvPrefix(utils.AppName)

	BindEnvs()

	viper.SetDefault(utils.Port, "8080")

	CheckMustBeSetEnvs()
}

func BindEnvs() {
	viper.BindEnv(utils.Port, "PORT")
	viper.BindEnv(utils.StorageAccessToken, "W3_STORAGE_TOKEN")
}

func EnvMustBeSet(key string) {
	if !viper.IsSet(key) {

		logger.Log.WithField(key, key).Fatalf("%v: not set", key)
	}
}

func CheckMustBeSetEnvs() {
	EnvMustBeSet(utils.StorageAccessToken)
}
