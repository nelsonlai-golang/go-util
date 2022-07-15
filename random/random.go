package random

import "math/rand"

var lowercase = []rune("abcdefghijklmnopqrstuvwxyz")
var uppercase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("0123456789")
var special = []rune("!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?`~")

type StringConfig struct {
	Lowercase bool
	Uppercase bool
	Numbers   bool
	Special   bool
}

func RandomString(length int, config StringConfig) string {
	var runes []rune
	if config.Lowercase {
		runes = append(runes, lowercase...)
	}
	if config.Uppercase {
		runes = append(runes, uppercase...)
	}
	if config.Numbers {
		runes = append(runes, numbers...)
	}
	if config.Special {
		runes = append(runes, special...)
	}
	b := make([]rune, length)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}
