package framework

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

// App is the entry point of the application, it can hold multiple server which implemented the Server interface
type App struct {
	logger  Logger
	servers []Server
}

func NewApp(logger Logger) *App {
	return &App{
		logger: logger,
	}
}

// WithServer will add the given server to the list of servers
func (app *App) WithServer(servers ...Server) {
	app.servers = append(app.servers, servers...)
}

// Start will start all the servers, it will block until the application is stopped
// this method will invoke the server's Serve method in a new goroutine
func (app *App) Start() error {
	eg, ctx := errgroup.WithContext(context.Background())
	for _, server := range app.servers {
		go func(server Server) {
			eg.Go(func() error {
				return server.Serve(ctx)
			})
		}(server)
	}
	eg.Go(func() error {
		return app.handleSysSignal(ctx)
	})
	return eg.Wait()
}

func (app *App) Logger() Logger {
	return app.logger
}

func (app *App) handleSysSignal(ctx context.Context) error {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	select {
	case <-ctx.Done():
		_ = app.logger.Log("msg", "exiting app...")
		return ctx.Err()
	case sig := <-sigs:
		_ = app.logger.Log("msg", fmt.Sprintf("receive sig %s, exiting...", sig))
		return fmt.Errorf("receive sig %v, exiting", sig)
	}
}
