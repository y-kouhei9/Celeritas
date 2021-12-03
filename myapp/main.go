package main

import (
	"github.com/y-kouhei9/celeritas"
	"myapp/handlers"
)

// application is the basic type of application
type application struct {
	App *celeritas.Celeritas
	Handlers *handlers.Handlers
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
