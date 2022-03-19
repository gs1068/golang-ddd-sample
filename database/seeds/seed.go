package main

import (
	"github.com/gs1068/golang_ddd_sample/config"
	seed "github.com/gs1068/golang_ddd_sample/database/seeds/seed_files"
)

func main() {
	db := config.NewDB()
	defer db.Close()
	seed.UsersSeed(db)
	seed.TasksSeed(db)
	seed.TimelinesSeed(db)
}
