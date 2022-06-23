package repository

import (
	"visitor-management-system/model"
)

func GetSubscriptionDetails(sub *model.Subscription) (*model.Subscription, error) {
	err := db.Where("company_id = ?", sub.CompanyId).Find(&sub).Error
	return sub, err
}
