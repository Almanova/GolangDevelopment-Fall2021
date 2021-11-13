package main

import (
	"context"
	"homeworks/w8/internal/http"
	"homeworks/w8/internal/store/postgres"
)

func main() {
	urlExample := "postgres://golang:golang@localhost:5432/skincare"
	store := postgres.NewDB()
	if err := store.Connect(urlExample); err != nil {
		panic(err)
	}
	defer store.Close()

	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		panic(err)
	}

	srv.WaitForGracefulTermination()
}
