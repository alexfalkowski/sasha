package robots

import "github.com/alexfalkowski/go-service/v2/net/http/mvc"

// Register robots.
func Register() {
	mvc.StaticFile("/robots.txt", "robots/robots.txt")
}
