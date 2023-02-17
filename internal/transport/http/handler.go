package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type CommentService interface{}

type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

func NewHandler(service CommentService) *Handler {
	handler := &Handler{
		Service: service,
	}

	handler.Router = mux.NewRouter()
	handler.mapRoutes()

	handler.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler.Router,
	}

	return handler
}

func (handler *Handler) mapRoutes() {
	handler.Router.HandleFunc("/hello", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(response, "Hello world!!!")
	})
}

func (handler *Handler) Serve() error {
	go func() {
		if err := handler.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	chanel := make(chan os.Signal, 1)
	signal.Notify(chanel, os.Interrupt)
	
	<-chanel

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	handler.Server.Shutdown(ctx)

	log.Println("shut down gracefully")
	return nil
}
