package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"githum.com/Vaiibhavv/students-api/students_api/internal/config"
)

func main() {

	// setup config
	cfg := config.MustLoad()

	//setup database
	//setup router
	router := http.NewServeMux()

	//seting up the url for response and request(get, post , method)
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to student api"))
	})

	// setup server , http.server is the struct
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	// this is the block then , it will not return anything ( we need to println before the server)
	slog.Info("server started", slog.String("address", cfg.HTTPServer.Address))
	//fmt.Printf("server started %s", cfg.HTTPServer.Address)

	// to handling the graceful shutdown(means interrupt while providing a api response)

	done := make(chan os.Signal, 1)

	// os signal notify to ther server server interrupt or not
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start the server")
		}
	}()

	<-done
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")

}
