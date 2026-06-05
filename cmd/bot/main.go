package main

import (
	"log"

	"cluster-mvp/internal/config"
	"cluster-mvp/internal/db"
	"cluster-mvp/internal/discord"
	"cluster-mvp/internal/services"
)

func main() {

	cfg := config.Load()

	db.InitDB()

	// Generate codes ONLY if database is empty
	services.SeedInviteCodes()

	discord.Start(cfg)

	log.Println("Cluster MVP is live")

	select {}
}