package robots

import "github.com/alexfalkowski/go-service/net/http/mvc"

// Register robots.
func Register() {
	mvc.Static("GET /robots.txt", "robots/robots.txt")
}
