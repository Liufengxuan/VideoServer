package dbops

import (
	_"github.com/go-sql-driver/mysql"

	"log"
	"database/sql"
)


func AddUserCredential(loginName string,pwd string)error{
	stmtIns,err:= dbConn.Prepare("insert into users (login_name,pwd) values (?,?)")
	defer stmtIns.Close()
	if err!=nil{
		log.Print("%s",err)
		return err
		}

		_,err=stmtIns.Exec(loginName,pwd)
		if err!=nil{
			  return err
		}

	return nil
}

func GetUserCredential(loginName string )(string,error)  {
	stmtOut,err:=dbConn.Prepare("select pwd from  users where login_name=?")
	if err!=nil{
		log.Print("%s",err)
		return "",err
	}

	var pwd string
	err=stmtOut.QueryRow(loginName).Scan(&pwd)
	if err!=nil&&err!=sql.ErrNoRows{
		return "",err
	}
	stmtOut.Close()

	return pwd,nil
}

func DeleteUser(loginName string,pwd string)error{
	stmtDel,err:=dbConn.Prepare("Delete from users where login_name=? and pwd=?")
	if err!=nil{
		log.Printf(" deleteUser err=%s",err)
		return err
	}
	_,err=stmtDel.Exec(loginName,pwd)
	if err!=nil&&err!=sql.ErrNoRows{
		return err
	}
	stmtDel.Close()
	return nil

}

func AddNewVideo(aid int,name string)(){

}