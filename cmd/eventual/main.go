package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/bldg14/eventual/internal/event"
	"github.com/bldg14/eventual/internal/event/stub"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/events", HandleGetAllEvents)

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	serverError := make(chan error, 1)
	go func() {
		log.Printf("listening on: %q\n", server.Addr)
		serverError <- server.ListenAndServe()
	}()

	select {
	case err := <-serverError:
		return fmt.Errorf("failed to ListenAndServe: %w", err)
	case <-ctx.Done():
		if err := server.Shutdown(context.Background()); err != nil {
			return fmt.Errorf("failed to Shutdown: %w", err)
		}
	}

	return nil
}

func HandleGetAllEvents(w http.ResponseWriter, r *http.Request) {
	var eventStoreStub stub.Stub
	events, err := event.GetAll(eventStoreStub)
	if err != nil {
		log.Printf("HandleGetAllEvents failed to GetAll: %s\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(events)
	if err != nil {
		log.Printf("HandleGetAllEvents failed to Marshal: %s\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
