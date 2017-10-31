package mysqldb

import(
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	lconf "github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
)
var GameDb sql.DB
func ConnectDB()  {
	sqlConnectStr :=lconf.MySqlUserName+":"+lconf.MySqlUserPwd+"@tcp("+lconf.MySqlAddr+")/"+lconf.DbName+"?charset=utf8"
	GameDb,err:=sql.Open("mysql",sqlConnectStr)
	if err!=nil{
		log.Error("failed to open database:",err.Error())
		return
	}
	defer GameDb.Close()
}

func CloseDB(){
	if GameDb!=nil{
		GameDb.Close()
	}
}

func InsertValue(sqlStr string)(id int64) {
	result, err := GameDb.Exec(sqlStr)
	id = -1
	if err != nil {
		log.Error("insert data failed:", err.Error())
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Error("fetch last insert id failed:", err.Error())
		return
	}
}

func UpdateValue(sqlStr string)(rowNum int64) {
	result, err := GameDb.Exec(sqlStr)
	rowNum = -1
	if err != nil {
		log.Error("update data failed:", err.Error())
		return
	}
	num, err := result.RowsAffected()
	rowNum = num
	if err != nil {
		log.Error("fetch row affected failed:", err.Error())
		return
	}
}

func DeleteValue(sqlStr string)(rowNum int64)  {
	result,err:=GameDb.Exec(sqlStr)
	rowNum=-1
	if err!=nil{
		log.Error("delete data failed:",err.Error())
		return
	}
	num,err:=result.RowsAffected()
	rowNum = num
	if err!=nil{
		log.Error("fetch row affected failed:",err.Error())
		return
	}
}