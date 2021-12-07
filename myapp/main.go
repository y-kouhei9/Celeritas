package main

import (
	"myapp/data"
	"myapp/handlers"

	"github.com/y-kouhei9/celeritas"
)

// application is the basic type of application
type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
