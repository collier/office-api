package sqldb

// Client contains all information about clients stored in the database
type Client struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	ProjectCount int     `json:"projectCount"`
	IsActive     bool    `json:"isActive"`
}

// GetAllClients queries the database, and returns a []Client for each row in
// the client table
func GetAllClients() ([]*Client, error) {
	rows, err := db.Query("select * from client")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	clients := make([]*Client, 0)
	for rows.Next() {
		var isActive int
		c := new(Client)
		err := rows.Scan(&c.ID, &c.Name, &c.Latitude, &c.Longitude, &c.ProjectCount, &isActive)
		if err != nil {
			return nil, err
		}
		c.IsActive = isActive != 0
		clients = append(clients, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return clients, nil
}
