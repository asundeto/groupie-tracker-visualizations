package yinyang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Parse() []Artist {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	netClient := http.Client{
		Timeout: time.Second * 10,
	}
	res, err := netClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	artists := []Artist{}
	jsonErr := json.Unmarshal(body, &artists)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	allRelations := RelatParse()
	for i := 0; i < len(artists); i++ {
		artists[i].DatesLocations = allRelations.IndexRels[i]
		artists[i].FirstAlbum = strings.ReplaceAll(artists[i].FirstAlbum, "-", ".")
	}
	for i := 0; i < len(artists); i++ {
		userData, err := json.Marshal(artists[i].DatesLocations)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		artists[i].DatesLocat = ChangeDateLocation(string(userData))
	}
	return artists
}

func RelatParse() Relat {
	var relations Relat

	parseRels, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return relations
	}
	defer parseRels.Body.Close()
	body, err := ioutil.ReadAll(parseRels.Body)
	if err != nil {
		return relations
	}

	err = json.Unmarshal(body, &relations)
	if err != nil {
		fmt.Println(err)
		return relations
	}
	return relations
}
