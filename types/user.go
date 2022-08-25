package types

import (
	"time"
	"visitor-management-system/model"
)

type User struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	BranchName string `json:"branch_name"`
	Name       string `json:"name"`
	BranchId   []int  `json:"branch_id"`
}

type Password struct {
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type Token struct {
	User_Token        string
	User_Refreshtoken string
	Branch            *[]model.Branch
}

type UserDetails struct {
	Id         int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name       string `json:"name" validate:"required,min=2,max=30"`
	Email      string `json:"email" validate:"required,email"`
	SubDomain  string
	CompanyId  int
	UserType   string `json:"user_type" validate:"eq=Admin|eq=Official"`
	BranchName string `json:"branch_name" validate:"required"`
	//Address    string    `json:"address"`
}

type BranchIds struct {
	BranchId []int64 `json:"branch_ids"`
}

type AllUser struct {
	Id              int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name            string `json:"name" validate:"required,min=2,max=30"`
	Email           string `json:"email" validate:"required,email"`
	SubDomain       string
	CompanyId       int
	BranchId        int
	UserType        string `json:"user_type" validate:"eq=Admin|eq=Official"`
	Password        string
	AllocatedBranch []Branches
}

type Branches struct {
	Id         int
	CompanyId  int
	BranchName string
	Address    string
}

type PaginationGetAllVisitor struct {
	TotalCount int64
	Items      []*model.Visitor
}

type PaginationGetAllRecord struct {
	TotalCount int64
	Items      []*model.Record
}

type Master struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Mastertoken struct {
	User_Token        string
	User_Refreshtoken string
}

type PackageDetails struct {
	PackageType string `json:"package_type" validate:"required"`
	Days        string `json:"days" validate:"required"`
}

type PaginationGetAllCompany struct {
	TotalCount int64
	Items      []*CompanyDetails
}

type PaginationGetAllPackage struct {
	TotalCount int64
	Items      []*model.Package
}

type PaginationGetAllAdmins struct {
	TotalCount int64
	Items      []*AdminDetails
}

type AdminDetails struct {
	Id          int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `json:"name" `
	Email       string `json:"email" `
	SubDomain   string
	CompanyId   int
	BranchId    int
	UserType    string `json:"user_type"`
	Password    string
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
}

type Status struct {
	Id     int  `json:"company_id"`
	Status bool `json:"status"`
}

type CompanyDetails struct {
	Id                 int    `gorm:"primary_key;AUTO_INCREMENT"`
	CompanyName        string `json:"company_name" validate:"required,min=2,max=30"`
	Address            string `json:"address"`
	SubDomain          string `json:"sub_domain"`
	SubscriberName     string `json:"subscriber_name"`
	SubscriberEmail    string `json:"subscriber_email" validate:"required,email"`
	Status             bool
	Subscription_Start time.Time
	Subscription_End   time.Time
	Package_Id         int `json:"package_id" validate:"required"`
	Subscription_type  string
}

type ChangeSubscription struct {
	CompanyId int `json:"company_id" validate:"required"`
	PackageId int `json:"package_id" validate:"required"`
}

type DataCount struct {
	TotalRegisteredVisitor int64 `json:"total_registered_visitor"`
	TodaysVisitor          int64 `json:"todays_visitor"`
	LastWeekVisitor        int64 `json:"last_week_visitor"`
	CurrentMonth           int64 `json:"current_month"`
	CurrentYear            int64 `json:"current_year"`
	LastMonthVisitor       int64 `json:"last_month_visitor"`
	LastYearVisitor        int64 `json:"last_year_visitor"`
	YesterDayVisitor       int64 `json:"yesterday_visitor"`
}

type AdminDataCount struct {
	TotalSubscriber    int64 `json:"total_subscriber"`
	ActiveSubscriber   int64 `json:"active_subscriber"`
	InactiveSubscriber int64 `json:"inactive_subscriber"`
	Package            []*AdmindataPackage
}

type AdmindataPackage struct {
	Package_type string `json:"package_type"`
	Count        int64  `json:"count"`
}

type VisitorRegistration struct {
	Message string
	Item    *model.Visitor
}
