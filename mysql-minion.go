package main

import (
	"fmt"
	"mysql-minion/mysql"
	"mysql-minion/utils"
)

var mysqlDumpBin string

func main() {

	creds := &utils.Creds{User: "play_mediadb", Password: "play_mediadb", Host: "localhost"}
	fmt.Printf("creds: %#v", creds)
	mysql.ShowTables(creds, "play_mediadb")
}
