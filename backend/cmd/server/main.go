package main

import (
	"github.com/dunstorm/sveltekit-golang-starter-template/backend/internal/app"
	"github.com/dunstorm/sveltekit-golang-starter-template/backend/internal/config"
)

// @title           Todo API
// @version         1.0
// @description     A simple todo list API
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /

func main() {
	cfg := config.New()

	app.CreateAndRun(cfg)
}
