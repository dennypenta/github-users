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
	// probably there is need to use godotenv lib and map in into struct,
	// so we could init it clearly and declare default values right in that struct
	limit, _ := strconv.Atoi(os.Getenv("MAX_PARALLEL_REQUESTS"))
	if limit < 1 {
		limit = defaultConcurrencyLimit
	}

	Concurrency.Limit = limit
}
