package main

import (
	"context"
	"fmt"

	db "github.com/kebyavonatlus/gorestapicourse/internal/database"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("successfully connected and pinged database")

	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
