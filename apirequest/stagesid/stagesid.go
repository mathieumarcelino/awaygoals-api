package stagesid

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"

	"awaygoals/clientcredentials"
)

type AllStages struct {
	Data []Data
}

type Data struct {
	Id int
	IdSeason int
	Name string
	hasStanding bool
}

type Id struct {
	Id8 int
	Id4 int
	Id2 int
	Id1 int
}


func StagesID(seasonsID int) {
	url := "https://football.elenasport.io/v2/seasons/"+ strconv.Itoa(seasonsID) +"/stages"
	fmt.Println(url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", clientcredentials.Clientcredentials())

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)


	var allStages AllStages
	json.Unmarshal([]byte(string(body)), &allStages)

	id8 := 0
	id4 := 0
	id2 := 0
	id1 := 0

	for i := 0; i < len(allStages.Data); i++ {
		if allStages.Data[i].Name == "8th Finals" {
			id8 = allStages.Data[i].Id
		}
		if allStages.Data[i].Name == "Quarter-finals" {
			id4 = allStages.Data[i].Id
		}
		if allStages.Data[i].Name == "Semi-finals" {
			id2 = allStages.Data[i].Id
		}
		if allStages.Data[i].Name == "Final" {
			id1 = allStages.Data[i].Id
		}
	}

	id := Id{
		Id8:	id8,
		Id4:	id4,
		Id2:	id2,
		Id1:	id1,
	}

	fmt.Println(id)
}