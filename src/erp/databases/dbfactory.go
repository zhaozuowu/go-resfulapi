package databases

import (
	"github.com/joho/godotenv"
	"fmt"
	"os"
)

type DbFactory struct {
}

func (dbFactory *DbFactory) GetDbDriver() Db {

	if err := godotenv.Load(".env"); err != nil {

		fmt.Printf("加载配置文件有误\n")

		os.Exit(1)
	}

	dbDriver := os.Getenv("DB_CONNECT")

	var dbObj Db

	switch dbDriver {

		case "mysql":
			dbObj = NewMysql()
		break;
		default:
			dbObj = NewMysql()

		break;

	}

	return dbObj
}
