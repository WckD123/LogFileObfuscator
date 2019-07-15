package models

import "github.com/jinzhu/gorm"

// DECLARING MODELS

type Logfile struct {
	gorm.Model
	Name         string   		`json:"name"`
	Age          int 			`json:"age"`
	Email        string 		`json:"email"`
	Mobile		 string 		`json:"mobile"`
	Role         string 		`json:"role"`
	Num          int  			`json:"num"`
	Address      string 		`json:"address"`
}


