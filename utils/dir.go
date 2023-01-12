package utils

import (
	"os"
)

func CreateDir() error {
    _, err := os.Stat("db")
    if os.IsNotExist(err) {
        err := os.Mkdir("./db", os.ModePerm)
        if err != nil {
            return err
        }
        return err
    }
    return nil
}
