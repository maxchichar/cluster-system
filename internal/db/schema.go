package db

import "log"

func InitSchema() {

	// invite codes table
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS invite_codes (
		code TEXT PRIMARY KEY,
		house TEXT,
		table_id INTEGER,
		slot TEXT,
		used INTEGER DEFAULT 0,
		used_by TEXT
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// users table
	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		discord_id TEXT PRIMARY KEY,
		house TEXT,
		table_id INTEGER,
		slot TEXT,
		verified_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Schema initialized")
}
