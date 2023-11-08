package business

import (
	"fmt"
)

func (s *Service) Export() error {

	units, err := s.getUnits()
	if err != nil {
		return err
	}

	fmt.Println(units)

	ships, err := s.getShips()
	if err != nil {
		return err
	}

	fmt.Println(ships)

	// A - guild.GetPlayers
	// B - loop players -> external.GetPlayerInfo()
	// C - report -> loop playersInfo()
	//		should generate a csv

	players, err := s.svcLogray.GetPlayers()
	if err != nil {
		return err
	}

	for _, p := range players {
		fmt.Println(p.GameName)
		s.getPlayerInfo(p)
	}

	return nil
}
