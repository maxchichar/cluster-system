package services

import (
	"errors"

	"cluster-mvp/internal/db"
)

type Result struct {
	House   string
	TableID int
	Slot    string
}

func VerifyCode(code string, discordID string) (Result, error) {

	var r Result
	var used int

	err := db.DB.QueryRow(`
		SELECT house, table_id, slot, used
		FROM invite_codes
		WHERE code = ?
	`, code).Scan(&r.House, &r.TableID, &r.Slot, &used)

	if err != nil {
		return Result{}, errors.New("invalid code")
	}

	if used == 1 {
		return Result{}, errors.New("code already used")
	}

	db.DB.Exec(`
		UPDATE invite_codes
		SET used = 1, used_by = ?
		WHERE code = ?
	`, discordID, code)

	db.DB.Exec(`
		INSERT OR REPLACE INTO users
		(discord_id, house, table_id, slot)
		VALUES (?, ?, ?, ?)
	`, discordID, r.House, r.TableID, r.Slot)

	return r, nil
}