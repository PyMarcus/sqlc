package main

import "fmt"

func main() {
	// migrate -path db/migration -database "postgresql://myuser:secret@localhost:5432/db?sslmode=disable" -verbose up
	fmt.Println("Hello, world!")
}
