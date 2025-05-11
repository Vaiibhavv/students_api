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
	"githum.com/Vaiibhavv/students-api/students_api/internal/http/handlers/student"
	sqlite "githum.com/Vaiibhavv/students-api/students_api/internal/storage/sqllite"
)

func main() {

	// setup config
	cfg := config.MustLoad()

	//setup database
	database, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	slog.Info("storage intialize", slog.String("version", "1.0.0"), slog.String("env %s", cfg.Env))

	//setup router
	router := http.NewServeMux()

	//seting up the url for response and request(get, poast , method)
	// here we are implemening the Storage interface signature method in sqlite package ( for database)
	router.HandleFunc("POST /api/students", student.New(database))
	router.HandleFunc("GET /api/students/{id}", student.GetById(database))
	router.HandleFunc("PUT /api/students/{id}", student.UpdateStudentById(database))

	// setup server , http.server is the struct
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	// this is the block then , it will not return anything ( we need to println before the server)
	slog.Info("server started", slog.String("address", cfg.HTTPServer.Address))
	//fmt.Printf("server started %s", cfg.HTTPServer.Address)

	// to handling the graceful shutdown(means interrupt while providing a api response)

	// by using done channel we can handled the interruption
	done := make(chan os.Signal, 1)

	// os signal notify to ther server server interrupt or not
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// this will run concurrently
	go func() {

		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	<-done
	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")

}
