package dbops

import (

	"testing"
)

func clearTables(){
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func  TestMain(m *testing.M)  {
clearTables()
m.Run()
clearTables()

}

func TestUserWorkFlow(t *testing.T){
	t.Run("Add",testAddUser)
	t.Run("Get",testGetUser)
	t.Run("Del",testDeleteUser)
	t.Run("Reget",testRegetUser)
}

func testAddUser(t *testing.T){
 	err:=AddUserCredential("lfx","1113")
 	if err!=nil{
 		t.Errorf("AddUser is err =%v",err)
	}
}
func testGetUser(t *testing.T){
	pwd,err:=GetUserCredential("lfx")
	if err!=nil{

		t.Errorf("GetUser is err =%s",err)
	}
	if pwd!="1113"{
		t.Errorf("get user error")
	}
	t.Logf("GetUser pwd is %s",pwd)

}
func testDeleteUser(t *testing.T){
	err:=DeleteUser("lfx","1113")
	if err!=nil{
		t.Errorf("Error of regetUser:%v",err)
	}

}
func testRegetUser(t *testing.T){
	pwd,err:=GetUserCredential("lfx")
	if err!=nil{
		t.Errorf("Error of regetUser:%v",err)
	}
	if pwd!=""{
		t.Errorf("delete user test failed")
	}
}