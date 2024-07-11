package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

func main() {
	tl := new(Tool)
	db, errDb := tl.getDb("coin.db")
	if errDb != nil {
		fmt.Printf("init db error:%s", errDb.Error())
		return
	}
	tl.createTable(db, "coin_list")

}
//git log --oneline --decorate --graph --all

type Tool struct {
}

func (p *Tool) getDb(dbName string) (*gorm.DB, error) {
	//创建连接
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				Colorful: false,
				LogLevel: logger.Info,
			},
		),
	})
	if err != nil {
		msg := "err not nil :" + err.Error()
		fmt.Println(msg)
		return nil, err
	}
	return db, nil
}

func (p *Tool) checkHasTable(mainDb *gorm.DB, tableName string) (bool, error) {
	where := map[string]string{
		"type": "table",
		"name": tableName,
	}
	var cnt int64
	tx := mainDb.Table("sqlite_master").Where(where).Count(&cnt)
	if tx.Error != nil {
		fmt.Printf("Count table err: %s", tx.Error.Error())
		return false, tx.Error
	}
	if cnt > 0 {
		return true, nil
	}
	return false, nil
}

func (p *Tool) createTable(mainDb *gorm.DB, tableName string) {
	//建表
	createSql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		coin_name VARCHAR(64) DEFAULT '',
		coin_code VARCHAR(64) DEFAULT '',
		price_int INTEGER DEFAULT 0,
		price_scale INTEGER DEFAULT 0,
		price_str VARCHAR(64) DEFAULT '',
		which_date VARCHAR(32) DEFAULT '',
		catch_time VARCHAR(64) DEFAULT '',
		from_where VARCHAR(128) DEFAULT '',
		other TEXT
	)`, tableName)
	var precision int
	fmt.Printf("%d", precision)
	dbCreate := mainDb.Exec(createSql)
	if dbCreate.Error != nil {
		fmt.Println("db create table failed:", dbCreate.Error.Error())
	}
	fmt.Println("exec create sql finish...")
}