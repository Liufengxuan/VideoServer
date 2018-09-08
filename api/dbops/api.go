package dbops

import (
	_"github.com/go-sql-driver/mysql"

	"log"
)




func AddUserCredential(loginName string,pwd string)error{
	stmtIns,err:= dbConn.Prepare("insert into users (login_name,pwd) values (?,?)")
	if err!=nil{
		return err
		}
		stmtIns.Exec(loginName,pwd)
	 stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string )(string,error)  {
	stmtOut,err:=dbConn.Prepare("select pwd from  users where login_name=?")
	if err!=nil{
		log.Print("s%",err)
		return "",err
	}

	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd,nil
}