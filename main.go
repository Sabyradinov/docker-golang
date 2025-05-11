package main

import (
	"fmt"
	"go-server/internal/http"
	"os"
	"os/signal"
	"syscall"

	"go-server/internal/config"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.Init("config.json")
	if err != nil {
		panic(fmt.Sprintf("config init error: %v", err))
	}

	gin, err := http.NewRouter(cfg)
	if err != nil {
		panic(fmt.Sprintf("gin router init error: %v", err))
	}

	server, err := http.New(cfg, gin)
	if err != nil {
		panic(fmt.Sprintf("create server error: %v", err))
	}

	startServerError := server.Start()

	var stopReason string
	select {
	case err = <-startServerError:
		stopReason = fmt.Sprintf("start server error: %v", err)
	case qs := <-quit:
		stopReason = fmt.Sprintf("received signal %s", qs.String())
	}

	fmt.Printf("%s\nshutting down server...\n", stopReason)
	err = server.Stop()
	if err != nil {
		fmt.Printf("stop server error: %v\n", err)
		return
	}

	fmt.Println("server stopped")
}
