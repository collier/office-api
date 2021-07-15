package sqldb

import (
	"database/sql"
)

// SnackCounter contains all snack counters stored in the database
type SnackCounter struct {
	ID               int64  `json:"id"`
	LaCroixCansDrunk string `json:"laCroixCansDrunk"`
}

// GetSnackCounters queries the database, and returns a SnackCounter struct
// which contains all the snack counters in the database
func GetSnackCounters() (*SnackCounter, error) {
	var sc SnackCounter
	err := db.QueryRow("select * from snack_counter limit ?", 1).Scan(&sc.ID, &sc.LaCroixCansDrunk)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &sc, nil
	}
}
