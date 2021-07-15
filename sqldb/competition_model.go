package sqldb

// Competition contains all information about competitons stored in the database
type Competition struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Winner      string `json:"winner"`
	Icon        string `json:"icon"`
	CompletedOn string `json:"completedOn"`
}

// GetCompetitions queries the database, and returns a []Competition for each
// competition, sorted by completed on date
func GetCompetitions() ([]*Competition, error) {
	rows, err := db.Query("select * from competition order by completed_on desc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	comps := make([]*Competition, 0)
	for rows.Next() {
		c := new(Competition)
		err := rows.Scan(&c.ID, &c.Name, &c.Winner, &c.Icon, &c.CompletedOn)
		if err != nil {
			return nil, err
		}
		comps = append(comps, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comps, nil
}

// DeleteCompetitionByID removes a row by id from the competition table in the
// database
func DeleteCompetitionByID(id int) error {
	stmt, err := db.Prepare("delete from competition where id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// AddCompetition adds a new row to the competition table in the database
func AddCompetition(c *Competition) error {
	q := `
		insert into competition 
		(name, winner, icon, completed_on)
		VALUES (?, ?, ?, ?)
	`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(c.Name, c.Winner, c.Icon, c.CompletedOn)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = id
	return nil
}

// UpdateCompetition updates a row in the competition table in the database
func UpdateCompetition(c *Competition) error {
	q := `
		update competition 
		set name=?, winner=?, icon=?, completed_on=?
		where id=?
	`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(c.Name, c.Winner, c.Icon, c.CompletedOn, c.ID)
	if err != nil {
		return err
	}
	return nil
}
