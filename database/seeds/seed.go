package main

import (
	"sync"

	"github.com/gs1068/golang_ddd_sample/config"
	seed "github.com/gs1068/golang_ddd_sample/database/seeds/seed_files"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	db := config.NewDB()
	defer db.Close()
	seed.UsersSeed(db)
	go seed.TasksSeed(db, &wg)
	go seed.TimelinesSeed(db, &wg)

	wg.Wait()
}
