package settings

import (
	"strconv"
	"os"
)

const (
	defaultConcurrencyLimit = 5
)

var Concurrency = struct {
	Limit int
} {}

func init() {
	limit, _ := strconv.Atoi(os.Getenv("MAX_PARALLEL_REQUESTS"))
	if limit < 1 {
		limit = defaultConcurrencyLimit
	}

	Concurrency.Limit = limit
}
