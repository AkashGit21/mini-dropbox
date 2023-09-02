package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AkashGit21/typeface-assignment/api"
	"github.com/AkashGit21/typeface-assignment/utils"
	"github.com/joho/godotenv"
)

// func main() {
// 	srv, err := NewServer()
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	StartServer(srv)
// 	// api, err := api.New()
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// // router.HandleFunc("/probes/readiness", func(res http.ResponseWriter, req *http.Request) {
// 	// // 	if err := db.PingContext(req.Context()); err != nil {
// 	// // 		res.WriteHeader(503)
// 	// // 	}
// 	// // })

// 	// srvHost := utils.GetEnvValue("APP_HOST", "localhost")
// 	// srvPort := utils.GetEnvValue("APP_PORT", "8081")
// 	// srv := &http.Server{
// 	// 	Addr:        fmt.Sprintf("%s:%v", srvHost, srvPort),
// 	// 	Handler:     router,
// 	// 	ReadTimeout: 2 * time.Minute,
// 	// }
// 	// // srv := &http.Server{
// 	// // 	Handler: router,
// 	// // 	Addr:    fmt.Sprintf(":%s", "8081"),
// 	// // }
// 	// go func() {
// 	// 	panic(srv.ListenAndServe())
// 	// }()

// 	// // Create channel for shutdown signals.
// 	// stop := make(chan os.Signal, 1)
// 	// signal.Notify(stop, os.Interrupt)
// 	// signal.Notify(stop, syscall.SIGTERM)

// 	// //Recieve shutdown signals.
// 	// <-stop
// 	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	// defer cancel()

// 	// if err := srv.Shutdown(ctx); err != nil {
// 	// 	log.Printf("error shutting down server %s", err)
// 	// } else {
// 	// 	log.Println("Server gracefully stopped")
// 	// }
// }

func NewServer() (*http.Server, error) {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	api, err := api.New()
	if err != nil {
		return nil, err
	}

	srvHost := utils.GetEnvValue("APP_HOST", "localhost")
	srvPort := utils.GetEnvValue("APP_PORT", "8081")
	srvAddress := fmt.Sprintf("%s:%v", srvHost, srvPort)
	log.Println("Configuring Server at address ", srvAddress)
	srv := http.Server{
		Addr:    srvAddress,
		Handler: api,
		// Read will Timeout after 2s if anything goes wrong.
		ReadTimeout: time.Duration(2 * time.Second),
	}

	return &srv, nil
}

func StartServer(srv *http.Server) {
	log.Println("Starting Server...")

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		log.Println("Shutting down the server gracefully...")
		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Println("HTTP server Shutdown: ", err)
			return
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Println("HTTP server ListenAndServe: ", err)
		return
	}

	<-idleConnsClosed
}
