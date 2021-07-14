package common

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var numRunes = []rune("0123456789")
var alphaRunes = []rune("abcdefghijklmnopqrstuvwxyz")
var alphaNumRunes = []rune(fmt.Sprintf("%s%s", string(numRunes), string(alphaRunes)))
var alphaNumWithUpperRunes = []rune(fmt.Sprintf("%s%s", string(alphaNumRunes), strings.ToUpper(string(alphaNumRunes))))

func GetRandString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 40)
	for i := range b {
		b[i] = alphaNumWithUpperRunes[rand.Intn(len(alphaNumWithUpperRunes))]
	}
	return string(b)
}
