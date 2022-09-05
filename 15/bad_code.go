package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var justString string

// createHugeString - creates huge string from radom chars
func createHugeString(size int) string {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	hugeString := strings.Builder{}
	for i := 0; i < size; i++ {
		hugeString.WriteRune('a' + rune(generator.Intn('z'-'a'+1)))
	}
	return hugeString.String()
}

func SomeFunc() {
	v := createHugeString(1 << 10)

	// justString = v[:100]

	// trouble with non-UTF-8 characters so better conver to rune
	justString = string([]rune(v)[:100])


	fmt.Println("justString = ", justString)


}

func main() {
	SomeFunc()
}
