package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"goschool/config"
)

var (
	Db       *gorm.DB
	ip       = config.Conf.Mysql.Ip
	port     = config.Conf.Mysql.Port
	mongoDb  = config.Conf.Mysql.Database
	user     = config.Conf.Mysql.User
	password = config.Conf.Mysql.Password
)

func mysqlUrl() (url string) {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, ip, port, mongoDb)
}
func InitDb()  {
	mysqlStr:= mysqlUrl()
	mysqlStr = mysqlStr + "?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(mysqlStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "t_",   // table name prefix, table for `User` would be `t_users`
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			//NameReplacer: strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})
	if err != nil {
		panic("connect mysql error")
	}

	if err != nil {
		panic(err)
	}
	fmt.Println("mysql connect success:", mysqlStr)

	Db = db
}
//func migration() {
//	// 自动迁移模式
//	db.AutoMigrate(&Admin{}, &Question{})
//}

//func InsertRowDemo(name, pw string) {
//	sqlStr := "insert into admin(nickname, password) values (?,?)"
//	ret, err := db.Exec(sqlStr, name, pw)
//	if err != nil {
//		fmt.Printf("insert failed, err:%v\n", err)
//		return
//	}
//	theID, err := ret.LastInsertId() // 新插入数据的id
//	if err != nil {
//		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("insert success, the id is %d.\n", theID)
//}
