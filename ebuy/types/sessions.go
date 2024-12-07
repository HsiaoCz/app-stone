package types

import "gorm.io/gorm"

type Sessions struct {
	gorm.Model
	Token string
	
}
