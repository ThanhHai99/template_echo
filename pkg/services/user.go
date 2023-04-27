package services

import "Template_Echo/pkg/models"

func CreateUser() string {
	return "Hello World"
}

func ReadAllUser() []*models.UserModel {
	var users []*models.UserModel
	DB().Model(&models.UserModel{}).Find(&users)
	return users
}

func ReadOneUser() string {
	return "Hello World"
}

func UpdateUser() string {
	return "Hello World"
}

func DeleteUser() string {
	return "Hello World"
}
