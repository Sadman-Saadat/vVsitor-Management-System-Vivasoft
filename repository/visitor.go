package repository

import (
	"fmt"
	"time"
	"visitor-management-system/model"
)

func CreateVisitor(visitor *model.Visitor) error {
	err := db.Create(&visitor).Error
	return err
}

func GetAllVisitor(id int) (visitor []*model.Visitor, err error) {
	err = db.Where("company_id = ?", id).Preload("TrackVisitors").Find(&visitor).Error
	return
}

func GetVisitor(visitor *model.Visitor) (*model.Visitor, error) {
	err := db.Find(&visitor).Error
	return visitor, err
}

func GetVisitorDetails(visitor *model.Visitor, id int) (*model.Visitor, error) {
	err := db.Where("company_id = ? AND id = ?", id, visitor.Id).Preload("TrackVisitors").Find(&visitor).Error
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

func GetTodaysVisitor(id int) ([]*model.Visitor, error) {
	var visitor []*model.Visitor
	val := "Arrived"
	today := time.Now().Local().Format("2006-01-02")
	err := db.Joins("JOIN track_visitors ON track_visitors.v_id = visitors.id AND track_visitors.company_id =?  AND track_visitors.date = ? AND track_visitors.status = ?", id, today, val).Preload("TrackVisitors", "date = ? ", today).Find(&visitor).Error
	return visitor, err
}

func GetTrackDetails(visitor *model.Visitor) (model.TrackVisitor, error) {
	var track model.TrackVisitor
	today := time.Now().Local().Format("2006-01-02")
	err := db.Where(" date=? AND company_id = ? AND v_id=?", today, visitor.CompanyId, visitor.Id).Find(&track).Error
	return track, err

}

func CheckOut(visitor *model.Visitor, track model.TrackVisitor) error {

	today := time.Now().Local().Format("2006-01-02")

	err := db.Where("company_id = ? AND v_id =? AND date=?", visitor.CompanyId, visitor.Id, today).Save(&track).Error
	return err
}

func IsVistorRegistered(email string, id int) (bool, error) {
	var visitor []*model.Visitor
	var count int64
	err := db.Where("email= ? AND company_id = ?", email, id).Find(&visitor).Count(&count).Error
	if count != 0 {
		return false, err
	}
	return true, err

}
