package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type distance struct {
	DestinationAddresses []string `json:"destination_addresses"`
	OriginAddresses      []string `json:"origin_addresses"`
	Rows                 []struct {
		Elements []struct {
			Distance struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"distance"`
			Duration struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"duration"`
			Status string `json:"status"`
		} `json:"elements"`
	} `json:"rows"`
	Status string `json:"status"`
}

func distancematrix(origin, destination, mode, key string) {
	var dis distance
	quick_reply := []resp{}

	url := "https://maps.googleapis.com/maps/api/distancematrix/json?origins=" + origin + "&destinations=" + destination + "&mode=" + mode + "&key=" + key

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		defer response.Body.Close()
		htmlData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(string(htmlData))
		json.Unmarshal(htmlData, &dis)
		fmt.Println(dis.DestinationAddresses)
		fmt.Println(dis.OriginAddresses)
		fmt.Println(dis.Rows[0].Elements[0].Distance.Text)
		fmt.Println(dis.Rows[0].Elements[0].Duration.Text)
	}
}

func main() {
	key := "AIzaSyBl6fX5Nswq-2iR6g-XPfFdxovBn0nJiTc"
	origin := "siruseri"
	destination := "kelambakkam"
	// transit_mode :="walking"
	mode := "driving"
	distancematrix(origin, destination, mode, key)
}
