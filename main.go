package main

import (
	"log"

	"github.com/musaubrian/tinygo/model"
	"github.com/musaubrian/tinygo/utils"
)

func main() {
	err := utils.CreateDir()
	if err != nil {
		log.Fatal("Error creating dir: ", err)
	}
//    utils.BoxTui()
	model.SetupDb()
	model.TinyGo()
}
