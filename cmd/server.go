package assignment

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

func main() {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=%s port=%d",
		"postgres", "postgres", "postgres", "localhost", "disable", 5432),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer closeDB(db)

	router := http.NewServeMux()

	router.HandleFunc("/probes/readiness", func(res http.ResponseWriter, req *http.Request) {
		if err := db.PingContext(req.Context()); err != nil {
			res.WriteHeader(503)
		}
	})

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", "8081"),
	}
	go func() {
		panic(srv.ListenAndServe())
	}()

	// Create channel for shutdown signals.
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	//Recieve shutdown signals.
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("error shutting down server %s", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}

func closeDB(db io.Closer) {
	if err := db.Close(); err != nil {
		_ = log.Println(errors.Wrap(err, "err closing db connection"))
	} else {
		_ = log.Println("db connection gracefully closed")
	}
}
