package main

import (
	_ "database/sql"
	_ "encoding/json"
	_ "fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "io/ioutil"
	"resources"
	_ "time"
	"routes"
)


func init() {

	// CONNECT TO THE DB
	resources.ConnectDB()

	// AUTOMIGRATE THE SCHEMA
	resources.MigrateDB()
}


func main() {

	//Create routes
	routes.CreateRoutes()

}



