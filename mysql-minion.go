package main

import "fmt"
import "os"


func main() {
	keys :=[]string{"MYSQL_BIN", "MYSQL_DUMP_BIN"}

	for i := 0; i < len(keys); i += 1 {
		fmt.Println(keys[i], ":", os.Getenv(keys[i]))
	}
}

