package repository

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"visitor-management-system/model"
)

func CreateVisitor(visitor *model.Visitor) error {
	err := db.Create(&visitor).Error
	return err
}

func GetAllVisitor(sql string) (visitor []*model.Visitor, err error) {
	err = db.Raw(sql).Scan(&visitor).Error
	return
}

func GetVisitor(visitor *model.Visitor) (*model.Visitor, error) {
	err := db.Find(&visitor).Error
	return visitor, err
}

func GetVisitorDetails(visitor *model.Visitor, id int) (*model.Visitor, error) {
	err := db.Where("company_id = ? AND id = ?", id, visitor.Id).Preload("TrackVisitors", func(db *gorm.DB) *gorm.DB {
		return db.Order("track_visitors.id DESC")
	}).Find(&visitor).Error
	return visitor, err
}

func UpdateVisitor(visitor *model.Visitor, id int) error {
	err := db.Where("company_id = ? AND id = ?", id, visitor.Id).Save(&visitor).Error
	return err
}

func Search(visitor *model.Visitor, id int) ([]*model.Visitor, error) {
	var list []*model.Visitor
	phone := fmt.Sprintf(visitor.Phone)
	phone += fmt.Sprintf("%s", "%")
	fmt.Println(phone)
	err := db.Where("phone LIKE ? AND company_id =?", phone, id).Find(&list).Error
	return list, err
}

func SearchForSpecificBranch(visitor *model.Visitor, company_id int, branch_id int) ([]*model.Visitor, error) {
	var list []*model.Visitor
	phone := fmt.Sprintf(visitor.Phone)
	phone += fmt.Sprintf("%s", "%")
	fmt.Println(phone)
	err := db.Where("phone LIKE ? AND company_id =? AND branch_id = ?", phone, company_id, branch_id).Find(&list).Error
	return list, err

}

func CheckIn(info *model.TrackVisitor) error {
	err := db.Create(&info).Error
	return err
}

func CountPresentVisitor(id int) (int64, error) {
	var count int64
	var count2 int64
	today := time.Now().Local().Format("2006-01-02")
	val := "Arrived"
	val2 := "left"
	var visitor []*model.TrackVisitor
	err := db.Where("status = ? AND date=? AND company_id = ?", val, today, id).Find(&visitor).Count(&count).Error
	err = db.Where("status = ? AND date=? AND company_id = ?", val2, today, id).Find(&visitor).Count(&count2).Error
	count = count + count2
	return count, err
}

func GetTodaysVisitor(sql string, startdate time.Time, enddate time.Time, status string) ([]*model.Record, error) {
	var visitor []*model.Record
	// val := "Arrived"
	// t := time.Now()
	// fmt.Println(t)
	// //t, err := time.Parse("2020-10-30 24:59:59", today)
	// // t := time.Date(today.Year, today.Month, today.Day, today.Hour, today.Minute, today.Second)
	// t2 := t.Add(time.Hour * time.Duration(24))
	if status != "" {
		err := db.Raw(sql, startdate, enddate, status).Scan(&visitor).Error
		return visitor, err
	}
	err := db.Raw(sql, startdate, enddate).Scan(&visitor).Error
	// err := db.Joins("JOIN track_visitors ON track_visitors.v_id = visitors.id AND track_visitors.company_id =?  AND track_visitors.date BETWEEN ? AND ? AND track_visitors.status = ?", id, t, t2, val).Preload("TrackVisitors", "date = ? ", t).Find(&visitor).Error

	return visitor, err
}

func GetTrackDetails(c_id int, id int) (model.TrackVisitor, error) {
	var track model.TrackVisitor
	times := time.Now().Local().Format("2006-01-02")
	fmt.Println(times)
	const shortForm = "2006-01-02"
	t, _ := time.Parse(shortForm, times)
	err := db.Where(" date=? AND company_id = ? AND id=?", t, c_id, id).Find(&track).Error
	return track, err

}

func CheckOut(id int, c_id int, track model.TrackVisitor) error {

	times := time.Now().Local().Format("2006-01-02")
	fmt.Println(times)
	const shortForm = "2006-01-02"
	today, _ := time.Parse(shortForm, times)

	err := db.Where("id =? AND company_id = ?  AND date=?", id, c_id, today).Save(&track).Error
	return err
}

func IsVistorRegistered(email string, id int, b_id int) (bool, error) {
	var visitor []*model.Visitor
	var count int64
	err := db.Where("email= ? AND company_id = ? AND branch_id = ?", email, id, b_id).Find(&visitor).Count(&count).Error
	if count != 0 {
		return false, err
	}
	return true, err

}
