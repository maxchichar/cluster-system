package services

import (
	"fmt"
	"math/rand"
	"time"

	"cluster-mvp/internal/db"
	"cluster-mvp/internal/houses"
)

var slots = []string{"S1", "S2", "S3", "S4"}

func SeedInviteCodes() {

	// Check if codes already exist
	var count int

	err := db.DB.QueryRow(
		"SELECT COUNT(*) FROM invite_codes",
	).Scan(&count)

	if err != nil {
		fmt.Println("Failed checking invite codes:", err)
		return
	}

	// If codes already exist, do nothing
	if count > 0 {
		fmt.Println("Invite codes already exist. Skipping generation.")
		return
	}

	rand.Seed(time.Now().UnixNano())

	for houseKey, h := range houses.Houses {
		for _, table := range h.Tables {
			for _, slot := range slots {

				code := generateCode(houseKey, table, slot)

				_, err := db.DB.Exec(`
					INSERT INTO invite_codes
					(code, house, table_id, slot, used)
					VALUES (?, ?, ?, ?, 0)
				`, code, houseKey, table, slot)

				if err != nil {
					fmt.Println("Insert error:", err)
					continue
				}

				fmt.Println("Generated:", code)
			}
		}
	}

	fmt.Println("Invite code generation complete.")
}

func generateCode(house string, table int, slot string) string {
	return fmt.Sprintf(
		"%s-T%02d-%s-%s",
		house,
		table,
		slot,
		randomString(6),
	)
}

func randomString(n int) string {

	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}