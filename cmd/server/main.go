package main

import (
	"context"
	"fmt"

	"github.com/kebyavonatlus/gorestapicourse/internal/comment"
	db "github.com/kebyavonatlus/gorestapicourse/internal/database"
	transportHttp "github.com/kebyavonatlus/gorestapicourse/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(commentService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
