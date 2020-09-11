package main

import (
	"github.com/joho/godotenv"
	"github.com/juniorrosul/delivery-much-challenge/adapters/primary"
)

func main() {
	godotenv.Load()
	primary.StartServer()
}
