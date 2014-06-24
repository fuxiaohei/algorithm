package trees

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
