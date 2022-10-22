package initializers

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

func ConnectToMysql(){
    // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
    dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("MYSQL_DATABASE"))

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil{
        log.Fatal(err)
    }

}
