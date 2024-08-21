package main

import (
	"gitea.koodsisu.fi/josepfrancescbenetmorella/literary-lions/dbaser"
)

func main() {
	dbaser.InitDb()
	dbaser.PopulateDb()
	// fmt.Println(dbaser.GetUser("jmadsen@uef.fi"))
}
