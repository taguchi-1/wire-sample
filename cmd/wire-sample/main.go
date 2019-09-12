package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/labstack/echo"
	"github.com/taguchi-1/wire-sample/infra/container"
)

const (
	version     = "v0.0.0"
	httpAddress = ":1323"
)

func main() {
	ctx := context.Background()

	print(ctx, "start http server\n")
	httpServer, err := newServer(ctx)
	if err != nil {
		print("failed to init server\n")
		os.Exit(1)
	}
	startServer(ctx, httpServer)

	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh

	print("shutdown http server\n")
	shutdownServer(ctx, httpServer)
}

func newServer(ctx context.Context) (*http.Server, error) {

	// Changed to wire generate code -
	// todoRepo := persistence.NewTodo()
	// todoService := service.NewTodo(todoRepo)
	// todoApp := application.NewTodo(todoService)
	// todoHandler := handler.NewTodo(todoApp)

	todoHandler, err := container.InitializeTodoHandler()
	if err != nil {
		return nil, err
	}

	e := echo.New()
	e.GET("/todos/:todoID", todoHandler.Get)
	return e.Server, nil
}

func startServer(ctx context.Context, httpServer *http.Server) {
	go func() {
		httpServer.Addr = httpAddress
		if err := httpServer.ListenAndServe(); err != nil {
			print("failed to start the http server\n")
		}
	}()
}

func shutdownServer(ctx context.Context, httpServer *http.Server) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := httpServer.Shutdown(ctx); err != nil {
			print("failed to shutdown the http server\n")
		}
		wg.Done()
	}()
	wg.Wait()
}
