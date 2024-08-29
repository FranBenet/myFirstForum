package main

import (
	"fmt"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
)

func main() {
	// dbaser.InitDb()
	// dbaser.PopulateDb()
	// user := models.User{"madrabbit@matrix.com", "whiterabbit", "Rz_;*$78)"}
	fmt.Println(dbaser.Posts())
}
