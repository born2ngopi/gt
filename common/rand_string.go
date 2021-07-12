package common

import "math/rand"

var letterRunes = []rune("0123456789abcdefghijKLMNOPQRSTUVWXYZ0123456789ABCDEFGHIJklmnopqrstuvwxyz")

func GetRandString() string {
	b := make([]rune, 40)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
