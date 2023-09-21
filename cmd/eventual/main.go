package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/kevinfalting/mux"
	"github.com/kevinfalting/structconf"

	"github.com/bldg14/eventual/internal/event"
	"github.com/bldg14/eventual/internal/event/stub"
	"github.com/bldg14/eventual/internal/middleware"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	flagEnv := flag.String("env", "local", "environment this server is running in")

	conf, err := structconf.New[config]()
	if err != nil {
		return fmt.Errorf("failed to structconf.New: %w", err)
	}
	flag.Parse()

	cfg := Config(*flagEnv)
	if err := conf.Parse(ctx, &cfg); err != nil {
		return fmt.Errorf("failed to Parse config: %w", err)
	}

	api := mux.New(
		middleware.CORS(strings.Split(cfg.AllowedOrigins, ",")...),
	)

	eh := mux.NewErrorHandler()

	api.Handle("/api/v1/events", mux.Methods(
		mux.WithGET(eh.Err(HandleGetAllEvents)),
	))

	server := http.Server{
		Handler: api,
		Addr:    fmt.Sprintf(":%d", cfg.Port),
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

func HandleGetAllEvents(w http.ResponseWriter, r *http.Request) error {
	var eventStoreStub stub.Stub
	events, err := event.GetAll(eventStoreStub)
	if err != nil {
		return mux.Error(fmt.Errorf("HandleGetAllEvents failed to GetAll: %w", err), http.StatusInternalServerError)
	}

	result, err := json.Marshal(events)
	if err != nil {
		return mux.Error(fmt.Errorf("HandleGetAllEvents failed to Marshal: %w", err), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

	return nil
}
