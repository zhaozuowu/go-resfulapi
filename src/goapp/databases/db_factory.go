package databases

type DbFactory struct {
}

func (dbObj DbFactory) GetConnect(driver string) Db {

	var db Db
	switch driver {
	case "mysql":
		db = Mysql{}
		break
	default:
		db = Mysql{}
		break

	}

	return db

}
