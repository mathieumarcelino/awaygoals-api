package generatejson

import (
    "fmt"

	seasonsid "awaygoals/apirequest/seasonsid"
    stagesid "awaygoals/apirequest/stagesid"
)

func Generatejson() {
    seasonsID := seasonsid.SeasonsID(2020, 2021)
    fmt.Println(seasonsID)

    stagesid.StagesID(seasonsID)
}