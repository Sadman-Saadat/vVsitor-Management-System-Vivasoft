package repository

import (
	"fmt"
	"gorm.io/gorm"
	"time"
	"visitor-management-system/model"
	//"visitor-management-system/types"
)

func CreateVisitor(visitor *model.Visitor) error {
	err := db.Create(&visitor).Error
	return err
}

func GetAllVisitor(sql string) (visitor []*model.Visitor, err error) {
	err = db.Raw(sql).Scan(&visitor).Error
	return
}

func GetAllVisitorSpecific(sql string, search string) (visitor []*model.Visitor, err error) {
	//search += fmt.Sprintf("%s", "%")
	if search != "" {
		err = db.Raw(sql, search, search, search).Scan(&visitor).Error
		return

	}
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

func Search(visitor *model.Visitor, id int, search string) ([]*model.Visitor, error) {
	var list []*model.Visitor
	search = fmt.Sprintf(search)
	search += fmt.Sprintf("%s", "%")
	fmt.Println(search)
	err := db.Where("phone LIKE ? OR name LIKE ? OR email LIKE ? AND company_id =?", search, search, search, id).Find(&list).Error
	return list, err
}

func SearchForSpecificBranch(visitor *model.Visitor, company_id int, branch_id int, search string) ([]*model.Visitor, error) {
	var list []*model.Visitor
	search = fmt.Sprintf(search)
	search += fmt.Sprintf("%s", "%")
	err := db.Where("company_id =? AND branch_id =? AND phone LIKE ? OR name LIKE ? OR email LIKE ?", company_id, branch_id, search, search, search).Find(&list).Error
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

func GetTodaysVisitor(company_id int, branch_id int, startdate time.Time, enddate time.Time, status string, search string, order string, offset int, limit int) ([]*model.Record, int64, error) {
	var visitor []*model.Record
	var count int64
	search += fmt.Sprintf("%s", "%")
	fmt.Println(limit)
	fmt.Println(offset)
	dbmodel := db.Model(&model.TrackVisitor{}).Select("track_visitors.*,visitors.name,visitors.email,visitors.phone,visitors.address,visitors.image_name,visitors.image_path,visitors.company_representating").Joins("left join visitors on visitors.id = track_visitors.v_id")
	dbmodel = dbmodel.Where("track_visitors.company_id = ? AND track_visitors.branch_id = ? AND track_visitors.date BETWEEN ? AND ?", company_id, branch_id, startdate, enddate)
	dbmodel = dbmodel.Where("track_visitors.status = ?", status)
	dbmodel = dbmodel.Where("visitors.email LIKE ? OR visitors.name LIKE ? OR visitors.phone LIKE ?", search, search, search)
	//dbmodel = dbmodel.Scan(&visitor)
	err := dbmodel.Count(&count).Error
	err = dbmodel.Limit(limit).Offset(offset).Scan(&visitor).Error

	return visitor, count, err
}

// // val := "Arrived"
// // t := time.Now()
// // fmt.Println(t)
// // //t, err := time.Parse("2020-10-30 24:59:59", today)
// // // t := time.Date(today.Year, today.Month, today.Day, today.Hour, today.Minute, today.Second)
// // t2 := t.Add(time.Hour * time.Duration(24))

// if status != "" {
// 	if search != "" {
// 		search += fmt.Sprintf("%s", "%")
// 		model := db.Raw(sql, startdate, enddate, status, search, search, search).Scan(&visitor)
// 		err := model.Error

// 		return visitor, err
// 	}
// 	model := db.Raw(sql, startdate, enddate, status).Scan(&visitor)
// 	err := model.Error

// 	return visitor, err
// }
// err := db.Raw(sql, startdate, enddate).Scan(&visitor).Error
// // err := db.Joins("JOIN track_visitors ON track_visitors.v_id = visitors.id AND track_visitors.company_id =?  AND track_visitors.date BETWEEN ? AND ? AND track_visitors.status = ?", id, t, t2, val).Preload("TrackVisitors", "date = ? ", t).Find(&visitor).Error

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

func CountVisitor(company_id int, id int, search string) (int64, error) {
	var visitor *model.Visitor
	var count int64

	if search != "" {
		search += fmt.Sprintf("%s", "%")
		err := db.Where("company_id = ? AND branch_id = ? AND name LIKE ? OR email LIKE ? OR phone LIKE ?", company_id, id, search, search, search).Find(&visitor).Count(&count).Error
		return count, err
	}
	err := db.Where("company_id = ? AND branch_id = ?", company_id, id).Find(&visitor).Count(&count).Error
	return count, err

}

func CountRecord(company_id int, branch_id int, status string, start time.Time, end time.Time, search string) (int64, error) {
	var visitor *model.TrackVisitor
	//var record []*model.Record
	var count int64
	//err := db.Raw(sql, start, end, status).Scan(&visitor).Count(&count).Error
	if status != "" {
		if search != "" {
			search += fmt.Sprintf("%s", "%")
			var count int64
			// res := db.Exec(`SELECT count(track_visitors.id) FROM track_visitors
			// LEFT JOIN visitors ON track_visitors.v_id = visitors.id
			//  WHERE (track_visitors.company_id = ? AND track_visitors.branch_id = ?
			//  AND track_visitors.status= ? AND track_visitors.date BETWEEN ? AND ?)
			//   AND (visitors.name LIKE ? OR visitors.email LIKE ? OR visitors.phone LIKE ?)`,
			//   company_id, branch_id, status, start, end, search, search, search)

			//   db.Model(&model.TrackVisitor{}).
			//   Count(&count).
			//   Joins("LEFT JOIN visitors ON track_visitors.v_id = visitors.id").
			//   Where(`track_visitors.company_id = ? AND track_visitors.branch_id = ?
			// 	AND track_visitors.status= ? AND track_visitors.date BETWEEN ? AND ?)
			// 	 AND (visitors.name LIKE ? OR visitors.email LIKE ? OR visitors.phone LIKE ?)`)
			//models.Exec
			// fmt.Println(res.Model(&model.TrackVisitor{}).Count(&count))
			// fmt.Println(count)

			return count, nil
		}

		err := db.Where("company_id = ? AND branch_id = ? AND status= ? AND date BETWEEN ? AND ?", company_id, branch_id, status, start, end).Find(&visitor).Count(&count).Error
		return count, err
	}

	err := db.Where("company_id = ? AND branch_id =? AND date BETWEEN ? AND ?", company_id, branch_id, start, end).Find(&visitor).Count(&count).Error

	return count, err

}
