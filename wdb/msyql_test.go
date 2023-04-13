package wdb

import (
	"fmt"
	"github.com/jeffcail/wildrocket/orm/models"
	"testing"
)

func TestInitGormMysql(t *testing.T) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/jiaxiao?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := InitGormMysql(dsn)
	if err != nil {
		fmt.Println(err)
	}
	admins := make([]*models.Admin, 0)
	db.Find(&admins)
	for _, v := range admins {
		fmt.Println(v)
	}
}
