package dbutils

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Scu struct {
	gorm.Model
	Scu_id     		uint64
	Sgu_id			uint64
	Name            	string	`gorm:"size:255"`
	Latitude		string	`gorm:"size:10"`
	Longitude		string	`gorm:"size:10"`
	Major			string	`gorm:"size:255"`
	Minor			string	`gorm:"size:255"`
	Is_connected		int	`gorm:"default:1"`
	Pwm             	int
	Status_updated_at  	time.Time
}

type Sgu struct {
	gorm.Model
	Sgu_id			uint64
	Name            	string	`gorm:"size:255"`
	Latitude		string	`gorm:"size:10"`
	Longitude		string	`gorm:"size:10"`
	Major			string	`gorm:"size:255"`
	Minor			string	`gorm:"size:255"`
	Is_connected		int	`gorm:"default:1"`
}