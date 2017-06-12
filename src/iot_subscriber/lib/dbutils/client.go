package dbutils

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var(
	DBCONN *gorm.DB
)
func InitDB(){
	DBCONN = newClient()
	DBCONN.AutoMigrate(&Sgu{}, &Scu{})
}

func newClient() *gorm.DB{
	db, err := gorm.Open("mysql", "HavellsDBAdmin:HavellsCCMS420@tcp(mysql.cqwf1pvghoch.us-west-2.rds.amazonaws.com:3306)/HavellsSubscriber?parseTime=True&loc=Local")
	if err != nil {
		revel.ERROR.Println(err.Error())
	}
	db.SetLogger(revel.WARN)
	db.LogMode(true)
	return db
}
