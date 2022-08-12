package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Company struct {
	Id                 int    `gorm:"primary_key;AUTO_INCREMENT"`
	CompanyName        string `json:"company_name" validate:"required,min=2,max=30"`
	Address            string `json:"address"`
	SubDomain          string `json:"sub_domain"`
	SubscriberName     string `json:"subscriber_name"`
	SubscriberEmail    string `json:"subscriber_email" validate:"required,email"`
	Status             bool
	Subscription_Start time.Time
	Subscription_End   time.Time
	Package_Id         int    `json:"package_id" validate:"required"`
	CompanySize        string `json:"company_size" validate:"required"`
	CompanyType        string `json:"company_type" validate:"required"`
	FaxNumber          string `json:"fax_no"`
	TotalBranchNumber  int    `json:"no_of_branch" validate:"required"`
	ContactNumber      string `json:"contact_num" validate:"required"`

	// Subscription    Subscription `gorm:"foreignKey:CompanyId;references:Id"`
	// User            []User       `gorm:"ForeignKey:CompanyId;references:Id"`

}
