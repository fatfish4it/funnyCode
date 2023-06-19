package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"math/rand"
	"os"
	"time"
)

type UserInfo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	Addr   string `json:"addr"`
	Other  string `json:"other"`
}

func main() {

	var user = new(UserInfo)
	//创建连接
	mainDb, errGet := user.getDb("user.db")
	if errGet != nil {
		fmt.Println("db init error:", errGet.Error())
		return
	}
	fmt.Printf("connect:%v\n", mainDb)

	//检查表是否存在
	tableName := "user_info"
	hasIn, errHas := user.checkHasTable(mainDb, tableName)
	if errHas != nil {
		fmt.Println("check table error:", errHas.Error())
		return
	}
	fmt.Printf("has table %s ? Say: %v\n", tableName, hasIn)
	if !hasIn {
		//创建表
		user.createTable(mainDb, tableName)
	}

	//查询
	//user.searchData(mainDb)

	//写入
	user.insertText(mainDb, tableName)

	//修改表
	//user.alterTable(mainDb)

	type T = struct{}
	type T2 struct{}
	var a T
	var a2 T2
	fmt.Printf("a:%v a2:%v", a, a2)

}

func (p *UserInfo) searchData(mainDb *gorm.DB) {
	/////查询方案-a
	//var tbList []UserInfo
	//mainDb.Raw("SELECT * FROM `text` WHERE id <= ?", 5).Scan(&tbList)

	//查询方案-b
	var tbList []UserInfo
	mainDb.Table("text").Select("id,name,url,remark").
		Where("id > ?", 5).
		Limit(3).
		Scan(&tbList)

	////查询方案-c
	//var tbList []UserInfo
	//mainDb.Table("text").Select("id,name,url,remark").Where(map[string]string{"id":"3"}).Scan(&tbList)
	for k, v := range tbList {
		fmt.Printf("k:%d v:%v\n", k, v)
	}
}

func (p *UserInfo) getDb(dbName string) (*gorm.DB, error) {
	////创建连接方案-a
	// db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	//创建连接方案-b
	mainDb, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{
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
	return mainDb, nil
}

func (p *UserInfo) insertText(mainDb *gorm.DB, tableName string) {
	var tbInsertData []UserInfo
	insertNum := 5
	//使每次随机的值都不相同，go-1.20之前需要
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < insertNum; i++ {
		randNum := rand.Float64()
		randMobile := rand.Int31()
		var tbOne = UserInfo{
			//兵者，诡道也。能而示之不能，近而示之远，强而示之弱。
			Name:   fmt.Sprintf("no.%d rand:%f", i, randNum),
			Mobile: fmt.Sprintf("65468131-%d", randMobile),
			Addr:   fmt.Sprintf("La.%f", randNum),
			Other:  fmt.Sprintf("insert at :%v", time.Now().Format("2006-01-02 15:04:05")),
		}
		tbInsertData = append(tbInsertData, tbOne)
	}
	//需要重新指定数据表，即调用 Table() 方法
	result := mainDb.Table(tableName).Create(&tbInsertData)
	if result.Error != nil {
		fmt.Printf("insert err not nil :%s\n", result.Error.Error())
	}
	fmt.Printf("row:%d\n", result.RowsAffected)
}

func (p *UserInfo) checkHasTable(mainDb *gorm.DB, tableName string) (bool, error) {
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

func (p *UserInfo) createTable(mainDb *gorm.DB, tableName string) {
	/*
	其他表
	CREATE TABLE "goods" (
		"id"	INTEGER UNIQUE,
		"name"	TEXT DEFAULT '',
		"price"	INTEGER DEFAULT 0,
		"from_where"	TEXT DEFAULT '',
		PRIMARY KEY("id" AUTOINCREMENT)
	);
	*/

	////建表方案-a
	//createSqlSlice := []string{
	//	"CREATE TABLE IF NOT EXISTS `user_info1` (",
	//	"`id` INTEGER,",
	//	"`name` VARCHAR(255),",
	//	"`other` TEXT",
	//	");",
	//}
	//createSql1 := strings.Join(createSqlSlice, "")
	//db1 := mainDb.Exec(createSql1)
	//if db1.Error != nil {
	//	fmt.Println("db1 create table failed:", db1.Error.Error())
	//}

	//建表方案-b
	createSql2 := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255),
		mobile VARCHAR(32) DEFAULT '',
		addr VARCHAR(32) DEFAULT '',
		other TEXT
	)`, tableName)
	db2 := mainDb.Exec(createSql2)
	if db2.Error != nil {
		fmt.Println("db2 create table failed:", db2.Error.Error())
	}
	fmt.Println("exec create sql finish...")
}

func (p *UserInfo) alterTable(mainDb *gorm.DB) {
	execSql := `ALTER TABLE	user_info ADD COLUMN addr VARCHAR(32) DEFAULT ''`
	db := mainDb.Exec(execSql)
	if db.Error != nil {
		fmt.Println("db alter table failed:", db.Error.Error())
	}
	fmt.Println("exec alter sql finish...")
}
