package main

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/piendop/postgresql/database"
)

var db *gorm.DB

func doCRUD() {
	//create a new employee
	dependents := json.RawMessage(`[{"relation":"friend", "info":{"first_name":"Ho", "last_name":"Phong"}},{"relation":"friend", "info":{"first_name":"Pham","last_name":"Chi"}}]`)
	hair := "black"
	nickname := "piendop"
	createdEmployee := database.Employee{
		FirstName:  "Do",
		LastName:   "Bao",
		Phones:     []string{"0777656331", "0787656331"},
		ManagerID:  nil,
		Dependents: postgres.Jsonb{dependents},
		Identity:   postgres.Hstore{"hair": &hair, "nickname": &nickname},
	}
	err := db.Debug().Create(&createdEmployee).Error
	if err != nil {
		fmt.Println("Cannot create a new employee", err.Error())
	} else {
		fmt.Println("Created Employee:", createdEmployee)
	}
	//read this employee (EmployeeID=1)
	readEmployee := database.Employee{
		EmployeeID: 1,
	}
	err = db.Debug().Model(&readEmployee).First(&readEmployee).Error
	if err != nil {
		fmt.Println("Cannot read employee with id=1", err.Error())
	} else {
		fmt.Println("Read Employee:", readEmployee)
		if err != nil {
			fmt.Println(err)
		} else {
			dependentsByte, _ := readEmployee.Dependents.MarshalJSON()
			var dependent []map[string]interface{}
			_ = json.Unmarshal(dependentsByte, &dependent)
			for _, val := range dependent {
				fmt.Println("name:", val["info"])
				fmt.Println("relation:", val["relation"])
			}
			fmt.Println("hair:", *readEmployee.Identity["hair"])
			fmt.Println("nickname:", *readEmployee.Identity["nickname"])
		}
	}

	//update info of this employee
	dependents = json.RawMessage(`[{"relation": "idol", "info":{"first_name":"john", "last_name":"cena"}}]`)
	hair = "still black"
	nickname = "bendoppler"
	updatedEmployee := database.Employee{
		EmployeeID: 1,
		Phones:     []string{"0777656331"},
		Dependents: postgres.Jsonb{dependents},
		Identity:   postgres.Hstore{"hair": &hair, "nickname": &nickname},
	}
	err = db.Debug().Model(&updatedEmployee).Update(&updatedEmployee).Error
	if err != nil {
		fmt.Println("Cannot update employee with id=1", err.Error())
	} else {
		fmt.Println("Updated Employee:", updatedEmployee)
		dependentsByte, _ := updatedEmployee.Dependents.MarshalJSON()
		var dependent []map[string]interface{}
		err = json.Unmarshal(dependentsByte, &dependent)
		for _, val := range dependent {
			info := val["info"].(map[string]interface{})
			fmt.Println("name:", info["first_name"], info["last_name"])
			fmt.Println("relation:", val["relation"])
			fmt.Println("hair:", *updatedEmployee.Identity["hair"])
			fmt.Println("nickname:", *updatedEmployee.Identity["nickname"])
		}
	}
	//delete info of this employee
	deletedEmployee := database.Employee{
		EmployeeID: 2,
	}

	err = db.Debug().Delete(&deletedEmployee).Error
	if err != nil {
		fmt.Println("Cannot delete employee with id=1", err.Error())
	} else {
		fmt.Println("Deleted Employee: ", deletedEmployee)
	}
}
func main() {
	//connect to postgres
	db = database.GetConnectionDb()
	//auto migrate: create table employee and update columns
	db.Debug().AutoMigrate(&database.Employee{})
	doCRUD()
	defer db.Close()
}
