package main

import (
	"fmt"

	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
)

func main() {
	// dbaser.InitDb()
	// dbaser.PopulateDb()
	fmt.Println(dbaser.Categories())
}
