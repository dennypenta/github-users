package input

import (
	"os"
	"strconv"
	"fmt"
)

func ParseArgs() []int64 {
	ids := make([]int64, 0)

	for i := 1; i < len(os.Args); i++ {
		id, err := strconv.ParseUint(os.Args[i], 10, 64)
		if err != nil {
			fmt.Println("isn't correct value, skiped: ", os.Args[i])
			continue
		}

		ids = append(ids, int64(id))
	}

	return ids
}
