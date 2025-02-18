package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstvywxz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(RandRange(1, 10))
}

func RandomPassword() string {
	return RandomString(RandRange(1, 5))
}

func RandomEmail() string {
	endType := []string{"gmail.com", "kntu.ac.ir", "yahoo.com", "sut.ac.ir"}
	len := len(endType)

	return fmt.Sprintf("%s@%s", RandomString(5), endType[rand.Intn(len)])
}

func RandRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomTitle() string {
	return RandomString(RandRange(1, 10))
}

func RandomDescription() string {
	return RandomString(RandRange(1, 10))
}
