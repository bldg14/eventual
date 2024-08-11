package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"

	"github.com/bldg14/eventual/internal/shell/http"
	"github.com/bldg14/eventual/internal/shell/storage"

	"github.com/kevinfalting/structconf"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := run(ctx, os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, args []string) error {
	fset := flag.NewFlagSet("eventual", flag.ExitOnError)
	flagEnv := fset.String("env", EnvLocal, "environment this server is running in")
	if err := fset.Parse(args); err != nil {
		return fmt.Errorf("failed to Parse flags: %w", err)
	}

	cfg := Config(*flagEnv)
	if err := structconf.Parse(ctx, &cfg); err != nil {
		return fmt.Errorf("failed to Parse config: %w", err)
	}

	dbURL, err := url.Parse(cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to Parse DatabaseURL: %w", err)
	}

	pool, err := storage.NewPool(ctx, storage.Config{
		DatabaseURL: dbURL,
	})
	if err != nil {
		return fmt.Errorf("failed to NewPool: %w", err)
	}

	server, err := http.NewServer(http.Config{
		Port:           cfg.Port,
		AllowedOrigins: strings.Split(cfg.AllowedOrigins, ","),
		Pool:           pool,
	})
	if err != nil {
		return fmt.Errorf("failed to NewServer: %w", err)
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
