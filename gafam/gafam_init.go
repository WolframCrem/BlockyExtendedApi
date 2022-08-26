package gafam

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var Gafam = loadGafamData()

const DataUrl string = "https://raw.githubusercontent.com/WolframCrem/BlockyExtendedApi/main/metadata/gafan.json"

func loadGafamData() []GafamData {
	response, err := http.Get(DataUrl)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(response.Body)
	var data []GafamData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
