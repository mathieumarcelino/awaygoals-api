package seasonsid

import (
	"encoding/json"
	"net/http"
	"io/ioutil"

	"awaygoals/clientcredentials"
)

type AllSeasons struct {
	Data []Data
}

type Data struct {
	Id int
	IdLeague int
	LeagueName string
	Start int
	End int
}

func SeasonsID(yearstart int, yearend int) int {
    url := "https://football.elenasport.io/v2/leagues/7/seasons"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", clientcredentials.Clientcredentials())

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var allSeasons AllSeasons
	json.Unmarshal([]byte(string(body)), &allSeasons)

	id := 0

	for i := 0; i < len(allSeasons.Data); i++ {
		if allSeasons.Data[i].Start == yearstart && allSeasons.Data[i].End == yearend {
			id = allSeasons.Data[i].Id
		}
	}

	return id
}
