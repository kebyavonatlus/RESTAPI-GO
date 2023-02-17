package main

import (
	"context"
	"fmt"

	"github.com/kebyavonatlus/gorestapicourse/internal/comment"
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

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("successfully connected and pinged database")

	commentService := comment.NewService(db)

	commentService.PostComment(
		context.Background(),
		comment.Comment {
			ID: "6a2fb5ea-20ea-48ef-a0cf-e22c485c4c67",
			Slug: "manual-test",
			Author: "Elliot",
			Body: "Hello world",
		},
	)


	fmt.Println(commentService.GetComment(
		context.Background(),
		"670af4c8-cf2b-448a-9aec-c15ce3b5b94e",
	))

	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
