package business

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/remifabas/swgohapi/pkg/guild"
)

type PlayerInfo struct {
	UnitInfos []UnitInfo `json:"units"`
	//Mods      []string   `json:"mods"`
	//Datacrons []string   `json:"datacrons"`
	//PlayerData []string   `json:"data"`
}

type UnitInfo struct {
	Datas Datas `json:"data"`
}

type Datas struct {
	BaseId    string `json:"base_id"`
	Name      string `json:"name"`
	GearLevel int    `json:"gear_level"`
}

func (s *Service) getPlayerInfo(p guild.Player) error {

	urlRequest := "http://api.swgoh.gg/player/" + p.AllyCode
	request, err := http.NewRequest(http.MethodGet, urlRequest, nil)
	if err != nil {
		panic(err)
	}

	response, err := s.httpClient.Do(request)
	if err != nil {
		panic(err)
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var pInfos []PlayerInfo

	if err := json.Unmarshal(resBody, &pInfos); err != nil {
		fmt.Println("Error unmarshaling json:", err)
		return err
	}

	fmt.Println(pInfos)

	return nil
}
