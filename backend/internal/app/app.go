package app

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/dunstorm/sveltekit-golang-starter-template/backend/internal/config"
)

type engine struct {
	ctx    context.Context
	cancel context.CancelFunc

	config  *config.Config
	runners *sync.WaitGroup
}

// CreateAndRun initializes and starts the application
func CreateAndRun(config *config.Config) {
	e := &engine{config: config}
	e.setupLogging()
	e.run()
}

func (e *engine) setupLogging() {
	if e.config.IsDev() {
		log.SetOutput(os.Stdout)
		log.Println("Development mode enabled")
	}
}

func (e *engine) run() {
	e.runners = &sync.WaitGroup{}
	e.ctx, e.cancel = context.WithCancel(context.Background())
	defer e.cancel()

	e.startWebServer()

	e.runners.Wait()
}
