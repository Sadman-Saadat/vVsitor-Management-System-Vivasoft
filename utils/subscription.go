package utils

import (
	"fmt"
	"time"
	//"visitor-management-system/const"
	//"visitor-management-system/model"
	"visitor-management-system/repository"
)

func ValidateSubscription(id int) (bool, string, error, bool) {

	res, err := repository.GetCompanyById(id)
	if err != nil {
		return false, "", err, false
	}

	fmt.Println("company")
	fmt.Println(res)
	if res.Status != true {
		return false, "you are restricted plz contact vivasoft", err, false
	}

	now := time.Now().Local()
	if now.After(res.Subscription_End) {
		return false, "your subscription is over", err, false
	}

	features, err := repository.GetPackageFeature(res.Package_Id)
	if err != nil {
		return false, "", err, features.Image
	}
	fmt.Println("featurs")
	fmt.Println(features)

	count, err := repository.CountPresentVisitor(id)
	fmt.Println("present")
	fmt.Println(count)
	if err != nil {
		return false, "", err, features.Image
	}
	if int(count) > features.VisitorCountPerDay {
		return false, "per day count limit exceeded", err, features.Image
	}

	total_count, err := repository.GetAllVisitor(id)
	fmt.Println("total visitor")
	fmt.Println(total_count)
	if err != nil {
		return false, "", err, features.Image
	}
	if int(total_count) > features.MaxRegistredVisitorCount {
		return false, "max limit exceeded", err, features.Image
	}

	return true, "", err, features.Image
}
