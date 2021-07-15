package wiki

import (
	"github.com/collier/office-api/config"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type resposne struct {
	Query struct {
		Pages struct {
			Num1678 struct {
				Pageid    int    `json:"pageid"`
				Ns        int    `json:"ns"`
				Title     string `json:"title"`
				Revisions []struct {
					Contentformat string `json:"contentformat"`
					Contentmodel  string `json:"contentmodel"`
					Content       string `json:"*"`
				} `json:"revisions"`
			} `json:"1678"`
		} `json:"pages"`
	} `json:"query"`
}

// StaffMember stuff
type StaffMember struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	CorpExtension string    `json:"corpExtension"`
	CellPhone     string    `json:"cellPhone"`
	Email         string    `json:"email"`
	Birthday      string    `json:"birthday"`
	StartDate     time.Time `json:"startDate"`
	EndDate       string    `json:"endDate"`
}

// GetActiveStaff stuff
func GetActiveStaff(username string, password string) ([]StaffMember, error) {
	urlRaw := "%s/mediawiki/api.php?action=query&prop=revisions&rvprop=content&format=json&formatversion=2&titles=Corporate+Directory"
	url := fmt.Sprintf(urlRaw, config.CompanyWikiURL)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(username, password)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var rawResp resposne
	err = json.Unmarshal(body, &rawResp)
	if err != nil {
		return nil, err
	}
	content := rawResp.Query.Pages.Num1678.Revisions[0].Content
	rows := strings.Split(content, "\n| ")
	// Drops the first row of items in the grid, as it contains the headers
	rows = append(rows[:0], rows[1:]...)
	// var employees []Employee
	staff := make([]StaffMember, 0)
	for i := range rows {
		props := strings.Split(rows[i], "||")
		if i == len(rows)-1 {
			props[7] = strings.Replace(props[7], "|-}", "", -1)
		} else {
			props[7] = strings.Replace(props[7], "|-", "", -1)
		}
		for j := range props {
			props[j] = strings.TrimSpace(props[j])
		}
		startDate, _ := time.Parse("Jan 02, 2006", props[6])
		if err != nil {
			continue
		}
		sm := StaffMember{
			ID:            props[0],
			Name:          props[1],
			CorpExtension: props[2],
			CellPhone:     props[3],
			Email:         props[4],
			Birthday:      props[5],
			StartDate:     startDate,
			EndDate:       props[7],
		}
		if sm.EndDate == "" {
			staff = append(staff, sm)
		}
	}
	return staff, nil
}
