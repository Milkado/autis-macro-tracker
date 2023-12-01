package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Print("Hello, World!")

	r := gin.Default()

	log.Fatal(r.Run(":8080"))
}