package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var trDB *gorm.DB
var trerr error

func init() {
	trDB, trerr = gorm.Open(mysql.Open("p2p:p2pA!123@tcp(192.168.66.149:3306)/zjbxinsurance?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if trerr != nil {
		panic("failed to connect database")
	}
}

type accounts struct {
	Id      int
	Balance int
}

type transactions struct {
	Id              int
	From_account_id string
	To_account_id   string
	Amount          int
}

func main() {
	// db.AutoMigrate(&accounts{}, &transactions{})
	// db.Create(&accounts{Id: 1, Balance: 300})
	// db.Create(&accounts{Id: 2, Balance: 300})
	tx := trDB.Begin()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("事务回滚")
			tx.Rollback()
		}
	}()
	// var a int = 0
	account1 := accounts{}
	tx.Where("id = ?", 1).Find(&account1)
	if account1.Balance >= 100 {
		tx.Debug().Model(&account1).Update("balance", account1.Balance-100)
		// b := 1 / a
		// fmt.Println(b)
		accounts2 := accounts{}
		tx.Where("id = ?", 2).Find(&accounts2)
		tx.Debug().Model(&accounts2).Update("balance", accounts2.Balance+100)
		tx.Create(&transactions{From_account_id: "1", To_account_id: "2", Amount: 100})
		tx.Commit()
	} else {
		fmt.Println("余额不足")
		tx.Rollback()
		return
	}
}
