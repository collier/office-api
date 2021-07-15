package sqldb

import (
	"time"
)

// CompanyEvent contains all information about company events stored in the
// database
type CompanyEvent struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	EventDate       time.Time `json:"eventDate"`
	Location        *string   `json:"location"`
	Icon            string    `json:"icon"`
	DisplayTimeFlag bool      `json:"displayTimeFlag"`
}

// GetFutureCompanyEvents queries the database, and returns a []CompanyEvent for
// each row in the company_event table, where the eventdate is after the
// current date
func GetFutureCompanyEvents() ([]*CompanyEvent, error) {
	now := time.Now()
	// Below line can simulate a given day
	// now, _ := time.Parse("2006-01-02", "2018-06-16")
	dateString := now.Format("2006-01-02")
	rows, err := db.Query("select * from company_event where event_date >= ? order by event_date asc", dateString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := make([]*CompanyEvent, 0)
	for rows.Next() {
		var displayTimeFlag int
		ce := new(CompanyEvent)
		err := rows.Scan(&ce.ID, &ce.Name, &ce.EventDate, &ce.Location, &ce.Icon, &displayTimeFlag)
		if err != nil {
			return nil, err
		}
		ce.DisplayTimeFlag = displayTimeFlag != 0
		events = append(events, ce)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// DeleteCompanyEventByID removes a row by id from the company_event table in
// the database
func DeleteCompanyEventByID(id int) error {
	q := "delete from company_event where id=?"
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

// AddCompanyEvent adds a new row to the company_event table in the database
func AddCompanyEvent(ce *CompanyEvent) error {
	q := `
		insert into company_event 
		(name, event_date, location, icon, display_time_flag)
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(ce.Name, ce.EventDate, ce.Location, ce.Icon, ce.DisplayTimeFlag)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	ce.ID = id
	return nil
}

// UpdateCompanyEvent updates a row in the company_event table in the database
func UpdateCompanyEvent(ce *CompanyEvent) error {
	q := `
		update company_event 
		set name=?, event_date=?, location=?, icon=?, display_time_flag=?
		where id=?
	`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ce.Name, ce.EventDate, ce.Location, ce.Icon, ce.DisplayTimeFlag, ce.ID)
	if err != nil {
		return err
	}
	return nil
}
