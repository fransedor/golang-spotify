package helper

import "math/rand"

func GenerateRandomString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randString := make([]rune, length)
	for i := range randString {
		randString[i] = letters[rand.Intn(len(letters))]
	}
	return string(randString)
}
