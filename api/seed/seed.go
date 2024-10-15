package seed

import (
	"encoding/json"
	"fmt"
	user "github.com/Octek/resource-profile-management-backend.git/api/users"
	"io/ioutil"
	"os"
	// "github.com/jinzhu/gorm"
)

type Seed struct {
	Name string
	Run  func() error
}

type JsonData struct {
	Categories []user.UserCategory
}

func GetJSONFileData() JsonData {
	jsonFile, err := os.Open("seed_data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened seed_data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var jsonData JsonData
	json.Unmarshal(byteValue, &jsonData)
	return jsonData
}

func SeedData(userService user.UserService) {
	var jsonData = GetJSONFileData()
	_ = userService.CreateCategories(jsonData.Categories)
}
