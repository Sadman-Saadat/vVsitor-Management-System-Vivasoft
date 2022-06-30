package repository

import (
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

func Search(visitor *model.Visitor, id int) (*model.Visitor, error) {
	err := db.Where("phone = ? AND company_id =?", visitor.Phone, id).Find(&visitor).Error
	return visitor, err
}

func CheckIn(info *model.TrackVisitor) error {
	err := db.Create(&info).Error
	return err
}

func CountPresentVisitor(id int) (int, error) {
	var count int
	today := time.Now().Local().Format("2006-01-02")
	val := "Arrived"
	var visitor []*model.TrackVisitor
	err := db.Where("status = ? AND date=? AND company_id = ?", val, today, id).Find(&visitor).Count(&count).Error
	return count, err
}

func GetTodaysVisitor(id int) ([]*model.Visitor, error) {
	var visitor []*model.Visitor
	val := "Arrived"
	today := time.Now().Local().Format("2006-01-02")
	err := db.Joins("JOIN track_visitors ON track_visitors.v_id = visitors.id AND track_visitors.date = ? AND track_visitors.status = ?", today, val).Preload("TrackVisitors", "date = ?", today).Find(&visitor).Error
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
