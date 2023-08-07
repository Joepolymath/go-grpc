package main

import (
	"fmt"

	"github.com/Joepolymath/go-grpc/database"
)

func init() {
	fmt.Println("Hello from Init")
	database.DatabaseConnection()
}

func main() {
	fmt.Println("Hello from main")
}