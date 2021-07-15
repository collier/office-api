package sqldb

import (
	"database/sql"
	"time"
)

// PetOfTheMonth contains all information about pet of the month stored in the
// database
type PetOfTheMonth struct {
	ID          int64     `json:"id"`
	OwnerName   string    `json:"ownerName"`
	PetName     string    `json:"petName"`
	PetSpecies  string    `json:"petSpecies"`
	Description *string   `json:"description"`
	Picture     string    `json:"picture"`
	Month       time.Time `json:"month"`
}

// GetPetOfTheMonth queries the database, and returns a PetOfTheMonth struct
// which contains the details of the current month's pet.
func GetPetOfTheMonth() (*PetOfTheMonth, error) {
	var p PetOfTheMonth
	now := time.Now()
	// now, _ := time.Parse("2006-01-02", "2018-10-01")
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	month := firstOfMonth.Format("2006-01-02")
	err := db.QueryRow("select * from pet_of_the_month where month=? limit ?", month, 1).Scan(
		&p.ID,
		&p.OwnerName,
		&p.PetName,
		&p.PetSpecies,
		&p.Description,
		&p.Picture,
		&p.Month,
	)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &p, nil
	}
}

// GetAllPetsOfTheMonth queries the database, and returns a PetOfTheMonth for
// the latest 12 entries in the pet_of_the_month table
func GetAllPetsOfTheMonth() ([]*PetOfTheMonth, error) {
	rows, err := db.Query("select * from pet_of_the_month order by month desc limit 12")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	potms := make([]*PetOfTheMonth, 0)
	for rows.Next() {
		p := new(PetOfTheMonth)
		err := rows.Scan(
			&p.ID,
			&p.OwnerName,
			&p.PetName,
			&p.PetSpecies,
			&p.Description,
			&p.Picture,
			&p.Month,
		)
		if err != nil {
			return nil, err
		}
		potms = append(potms, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return potms, nil
}

// DeletePetOfTheMonthByID removes a row by id from the pet_of_the_month table
// in the database
func DeletePetOfTheMonthByID(id int) error {
	stmt, err := db.Prepare("delete from pet_of_the_month where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// AddPetOfTheMonth adds a new row to the pet_of_the_month table in the database
func AddPetOfTheMonth(p *PetOfTheMonth) error {
	q := `
		insert into pet_of_the_month 
		(owner_name, pet_name, pet_species, description, picture, month)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(
		p.OwnerName,
		p.PetName,
		p.PetSpecies,
		p.Description,
		p.Picture,
		p.Month,
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = id
	return nil
}

// UpdatePetOfTheMonth updates a row in the pet_of_the_month table in the database
func UpdatePetOfTheMonth(p *PetOfTheMonth) error {
	q := `
		update pet_of_the_month 
		set owner_name=?, pet_name=?, pet_species=?, description=?, picture=?, month=?
		where id=?
	`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		p.OwnerName,
		p.PetName,
		p.PetSpecies,
		p.Description,
		p.Picture,
		p.Month,
		p.ID,
	)
	if err != nil {
		return err
	}
	return nil
}
