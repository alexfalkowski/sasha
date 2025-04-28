package meta

import (
	"time"

	"github.com/alexfalkowski/go-service/env"
)

// NewYear for meta.
func NewYear() Year {
	return Year(time.Now().Year())
}

// Year is the current year we are in.
type Year int

// NewInfo for meta.
func NewInfo(version env.Version, year Year) *Info {
	return &Info{Version: version, Year: year}
}

// Info for all the models.
type Info struct {
	Version env.Version
	Year    Year
}
