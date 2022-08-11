package repository

import (
	"fmt"
	"time"
	"visitor-management-system/model"
	"visitor-management-system/types"
)

func CreateUser(user *model.User) (*model.User, error) {
	err := db.Create(&user).Error
	return user, err
}

func GetAllUsers(id int) ([]types.UserDetails, error) {
	join_sql := "SELECT users.id,users.name,users.email,users.sub_domain ,users.company_id,users.branch_id,users.user_type,branches.branch_name,branches.address FROM users LEFT JOIN branches ON users.branch_id = branches.id WHERE users.company_id = ?"
	var official_users []types.UserDetails
	//err := db.Where("company_id = ?", id).Find(&official_users).Error
	err := db.Raw(join_sql, id).Scan(&official_users).Error
	return official_users, err
}

func DeleteOfficialUser(id int) error {
	//sql := fmt.Sprintf("DELETE FROM users WHERE users.id = %d", id)
	var user model.User
	user.Id = id
	err := db.Delete(&user)
	fmt.Println(err.RowsAffected)
	return nil
}

func UpdateOfficialUser(user *model.User) error {
	err := db.Save(&user).Error
	return err
}

func GetUserByEmail(email string, subdomain string) (*model.User, error) {
	var user model.User

	err := db.Where("sub_domain = ? AND email = ?", subdomain, email).Find(&user).Error
	return &user, err
}

func GetBranchDetails(id int, bid int) (*model.Branch, error) {
	var branch model.Branch
	branch.Id = bid
	fmt.Println(bid)
	fmt.Println(branch)
	err := db.Find(&branch).Error
	return &branch, err
}

func GetUserById(id int) (*model.User, error) {
	var user *model.User

	err := db.Where("id= ?", id).Find(&user).Error
	return user, err
}

func GetData(data *types.DataCount, company_id int, branch_id int) (*types.DataCount, error) {
	var count int64
	const shortForm = "2006-01-02"

	err := db.Model(&model.Visitor{}).Where("company_id = ? AND branch_id= ?", company_id, branch_id).Count(&count).Error
	data.TotalRegisteredVisitor = count

	times := time.Now().Local().Format("2006-01-02")

	t, _ := time.Parse(shortForm, times)
	err = db.Model(&model.TrackVisitor{}).Where("company_id=? AND branch_id=? AND date =?", company_id, branch_id, t).Count(&count).Error
	data.TodaysVisitor = count

	start := time.Now().Local().Format("2006-01-02")
	t2, _ := time.Parse(shortForm, start)

	end := time.Now().Local().AddDate(0, 0, -1).Format("2006-01-02")
	t1, _ := time.Parse(shortForm, end)
	err = db.Model(&model.TrackVisitor{}).Where("company_id=? AND branch_id=? AND date = ?", company_id, branch_id, t1).Count(&count).Error
	data.YesterDayVisitor = count

	end = time.Now().Local().AddDate(0, 0, -7).Format("2006-01-02")
	t1, _ = time.Parse(shortForm, end)
	err = db.Model(&model.TrackVisitor{}).Where("company_id=? AND branch_id=? AND date BETWEEN ? AND ?", company_id, branch_id, t1, t2).Count(&count).Error
	data.LastWeekVisitor = count
	///////////////////////////////////////////////////////////////
	month := time.Now().Local().AddDate(0, 0, 0)
	monthstr := month.AddDate(0, -1, -month.Day()+1).Format("2006-01-02")
	t1, _ = time.Parse(shortForm, monthstr)
	month_end := t1.AddDate(0, 0, 31).Format("2006-01-02")
	t2, _ = time.Parse(shortForm, month_end)
	err = db.Model(&model.TrackVisitor{}).Where("company_id=? AND branch_id=? AND date BETWEEN ? AND ?", company_id, branch_id, t1, t2).Count(&count).Error
	data.LastMonthVisitor = count

	month = time.Now().Local().AddDate(0, 0, 0)
	monthstr = month.AddDate(0, 0, -month.Day()+1).Format("2006-01-02")
	t1, _ = time.Parse(shortForm, monthstr)
	month_end = time.Now().Local().AddDate(0, 0, 0).Format("2006-01-02")
	t2, _ = time.Parse(shortForm, month_end)
	err = db.Model(&model.TrackVisitor{}).Where("company_id=? AND branch_id=? AND date BETWEEN ? AND ?", company_id, branch_id, t1, t2).Count(&count).Error
	data.CurrentMonth = count

	month = time.Now().Local().AddDate(0, 0, 0)
	monthstr = month.AddDate(-1, 1-int(month.Month()), -month.Day()+1).Format("2006-01-02")
	t1, _ = time.Parse(shortForm, monthstr)
	month_end = t1.AddDate(0, 11, 31).Format("2006-01-02")
	t2, _ = time.Parse(shortForm, month_end)
	err = db.Model(&model.TrackVisitor{}).Where("company_id=? AND branch_id=? AND date BETWEEN ? AND ?", company_id, branch_id, t1, t2).Count(&count).Error
	data.LastYearVisitor = count

	month = time.Now().Local().AddDate(0, 0, 0)
	monthstr = month.AddDate(0, 1-int(month.Month()), -month.Day()+1).Format("2006-01-02")
	t1, _ = time.Parse(shortForm, monthstr)
	month_end = time.Now().Local().AddDate(0, 0, 0).Format("2006-01-02")
	t2, _ = time.Parse(shortForm, month_end)
	err = db.Model(&model.TrackVisitor{}).Where("company_id=? AND branch_id=? AND date BETWEEN ? AND ?", company_id, branch_id, t1, t2).Count(&count).Error
	data.CurrentYear = count
	return data, err
}
