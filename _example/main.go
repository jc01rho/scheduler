package main

import (
	"fmt"
	"log"
	"time"

	"github.com/rakanalh/scheduler"
	"github.com/rakanalh/scheduler/storage"
)

func Hello(name string) {
	fmt.Println("Hello", name)
}

func Recurring(name string) {
	fmt.Println("Hello from recurring", name)
}

func main() {
	storage := storage.NewSqlite3Storage(
		storage.Sqlite3Config{
			DbName: "db.store",
		},
	)
	if err := storage.Connect(); err != nil {
		log.Fatal("Could not connect to db", err)
	}

	if err := storage.Initialize(); err != nil {
		log.Fatal("Could not intialize database", err)
	}

	s := scheduler.New(storage)

	if err := s.RunAfter(5*time.Second, Hello, "Rakan"); err != nil {
		panic(err)
	}
	if err := s.RunEvery(5*time.Second, Recurring, "Rakan"); err != nil {
		panic(err)
	}
	s.Start()
	s.Wait()
}