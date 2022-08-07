package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type PackageFeatures struct {
	Id                       int  `gorm:"primary_key;AUTO_INCREMENT"`
	PackageId                int  `json:"package_id" validate:"required"`
	VisitorCountPerDay       int  `json:"visitor_count_per_day" validate:"required"`
	MaxRegistredVisitorCount int  `json:"max_registered_visitor_count" validate:"required"`
	Image                    bool `json:"image" validate:"required"`
	//MaxImageCount            int  `json:"max_image_count" validate:"required"`
	Email bool `json:"email" validate:"required"`
}
