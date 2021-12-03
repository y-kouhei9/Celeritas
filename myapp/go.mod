module myapp

go 1.17

replace github.com/y-kouhei9/celeritas => ../celeritas

require (
	github.com/go-chi/chi/v5 v5.0.7
	github.com/y-kouhei9/celeritas v0.0.0-00010101000000-000000000000
)

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v6 v6.1.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
