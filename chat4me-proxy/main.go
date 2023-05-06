package main

import (
	"log"
	"os"
)

var (
	cfg    c4mrConfig
	server c4mrServer
)

func main() {
	log.Println("Initializing Chat 4 Me request proxy")
	cfg = c4mrConfig{}
	var err error
	if cfg.workingDir, err = os.Getwd(); err != nil {
		log.Fatalln("Error getting working directory:", err.Error())
	}
	initConfig()
	initOpenAI()
	initServer(nil)
}
