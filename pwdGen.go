package main

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/f1bonacc1/glippy"
)

// Returns a random 15 character string from the collection
func generatePassword() string{

    var collection = []byte("0123456789qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
    pwdLength := 15
    pwd := make([]byte, pwdLength)
    rand.Read(pwd)

    for i := 0; i < pwdLength; i++ {
        pwd[i] = collection[int(pwd[i]%byte(len(collection)))]
    }

    return string(pwd)
}

// Copies generated string to clipboard
func CopyToClipboard() {
    text := generatePassword()
    err := glippy.Set(text)
    if err != nil {
        log.Fatal("Could not copy to clipboard", err)
    }
    fmt.Println(`
    *************************
       copied to clipboard
    `)
}
