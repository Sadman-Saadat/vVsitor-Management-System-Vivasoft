package repository

import (
	//"visitor-management-system/database"
	//"fmt"
	"visitor-management-system/model"
)

func CreatePackageFeatures(features *model.PackageFeatures) (*model.PackageFeatures, error) {
	err := db.Create(&features).Error
	return features, err
}

func GetPackageFeature(id int) (*model.PackageFeatures, error) {
	var feature *model.PackageFeatures
	err := db.Model(&model.PackageFeatures{}).Where("package_id =?", id).Find(&feature).Error
	return feature, err
}
