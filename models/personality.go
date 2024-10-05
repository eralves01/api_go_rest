package models

import (
	"api_go_rest/database"

	"gorm.io/gorm"
)

type Personality struct {
	Id      int    `json:"Id"`
	Name    string `json:"Name"`
	History string `json:"History"`
}

func GetAllPersonalities() []Personality {
	var personalities []Personality
	dbConnection := database.DBConnect()
	dbConnection.Find(&personalities)

	return personalities
}

func GetPersonalityById(id string) Personality {
	var personality Personality
	dbConnection := database.DBConnect()
	dbConnection.First(&personality, id)

	return personality
}

func CreatePersonality(personality Personality) *gorm.DB {
	dbConnection := database.DBConnect()
	result := dbConnection.Create(&personality)

	return result
}

func DeletePersonality(id string) *gorm.DB {
	var personality Personality
	dbConnection := database.DBConnect()
	result := dbConnection.Delete(&personality, id)

	return result
}

func PutPersonality(personality Personality) *gorm.DB {
	dbConnection := database.DBConnect()
	result := dbConnection.Save(&personality)

	return result
}
