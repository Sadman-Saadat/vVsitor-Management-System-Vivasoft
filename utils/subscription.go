package utils

import (
	"fmt"
	"time"
	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
)

func ValidateSubscription(id int) (bool, string, error) {
	var subscription = new(model.Subscription)
	subscription.CompanyId = id
	res, err := repository.GetSubscriptionDetails(subscription)
	if err != nil {
		return false, "", err
	}
	now := time.Now().Local()

	if now.After(res.Subscription_end) {
		return false, "your subscription is over", err
	}

	count, err := repository.CountPresentVisitor()
	fmt.Println(count)
	if err != nil {
		return false, "", err
	}
	if res.Subscription_type == "silver" && count > 10 {
		return false, consts.Upgrade, err
	}

	return true, "", err
}
