package database

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

type Employee struct {
	EmployeeID int32           `gorm:"AUTO_INCREMENT; primary_key; unique_index"`
	FirstName  string          `gorm:"type:varchar(50);not null"`
	LastName   string          `gorm:"type:varchar(50);not null"`
	ManagerID  *int32          `gorm:"type:int"`
	Phones     pq.StringArray  `gorm:"type:varchar(10)[]"`
	Dependents postgres.Jsonb  `gorm:"type:json"`
	Identity   postgres.Hstore `gorm:"type:hstore"`
}
