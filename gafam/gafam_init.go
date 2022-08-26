package gafam

import (
	"encoding/json"
	"io"
	"net/http"
)

var LoadedGafam = loadGafamData()

const DataUrl string = "https://raw.githubusercontent.com/WolframCrem/BlockyExtendedApi/main/metadata/gafan.json"

func loadGafamData() []GafamData {
	var data []GafamData
	response, err := http.Get(DataUrl)
	if err != nil {
		return data
	}
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		return data
	}
	return data
}
