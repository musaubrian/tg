package utils

import (
	"crypto/rand"

	"github.com/atotto/clipboard"
)

// Returns a random 15 character string from the collection
func GeneratePassword() string {
	collection := []byte("0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM~!@#$%^&*(()-?/")

	pwdLength := 15
	pwd := make([]byte, pwdLength)
	rand.Read(pwd)

	for i := 0; i < pwdLength; i++ {
		pwd[i] = collection[int(pwd[i]%byte(len(collection)))]
	}

	return string(pwd)
}

// Copies generated string to clipboard
func CopyToClipboard(text string) error {
	err := clipboard.WriteAll(text)
	return err
}
