module example.com/goapi

go 1.16

require (
	example.com/routes v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.2
)

replace example.com/routes => ./routes
