package main

import (
	"fmt"

	guild "github.com/remifabas/swgohapi/pkg/guild"
	b "github.com/remifabas/swgohapi/pkg/swgohapi/business"
)

func main() {
	fmt.Println("SWGOH API")

	guildSvc := guild.NewService("players.yaml")
	swSvc := b.NewService(b.SWgohConfig{
		URL: "http://api.swgoh.gg/player/"},
		guildSvc,
	)
	err := swSvc.Export()
	if err != nil {
		panic(err)
	}

	// c := http.Client{
	// 	Timeout: time.Duration(1) * time.Second,
	// }

	// request, err := http.NewRequest(http.MethodGet, "http://api.swgoh.gg/player/327786519", nil)
	// if err != nil {
	// 	panic(err)
	// }

	// response, err := c.Do(request)
	// if err != nil {
	// 	panic(err)
	// }

	// resBody, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(resBody))

}
