package dbops

import (

	"testing"
	"strconv"
	"time"
	"log"
)
var tempVid string

func clearTables(){
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	//dbConn.Exec("truncate comments")
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

func TestVideoWorkFlow(t *testing.T){
	clearTables()
	t.Run("PrepareUser",testAddUser)
	t.Run("AddVideo",testAddVideoInfo)
	t.Run("GetVideo",testGetVideoInfo)
	t.Run("DelVideo",testDeleteVideoInfo)
	t.Run("RegetVideo",testRegetVideo)

}

func TestComments(t *testing.T){
	clearTables()
	t.Run("PrepareUser",testAddUser)
	t.Run("AddComments",testAddComments)
	t.Run("GetComments",testListComments)
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

func testAddVideoInfo(t *testing.T) {
	vi,err:=AddNewVideo(1,"myvideo")
	if err!=nil{
		t.Errorf("Error of AddVideoInfo: %v",err)
	}
	tempVid=vi.Id

}
func testGetVideoInfo(t *testing.T){
	vi,err:=GetVideoInfo(tempVid)
	if err!=nil{
		t.Errorf("Error of GetVideoInfo:%v",err)
	}
	if vi==nil{
		t.Errorf("GetVideoInfo is err : %v",err)
	}
}
func testDeleteVideoInfo(t *testing.T){
	err:=DeleteVideoInfo(tempVid)
	if err!=nil{
		t.Errorf("Error of DeleteVideoInfo:%v",err)
	}

}
func testRegetVideo(t *testing.T){
	vi,err:=GetVideoInfo(tempVid)
	if err!=nil{
		t.Errorf("Error of RegetVideo:%v",err)
	}
	if vi!=nil{
		t.Errorf("RegetVideo  err ")
	}
}

func testAddComments(t *testing.T){
	vid:="123456"
	aid:=1
	content:="i like this video"

	err=AddNewComments(vid,aid,content)

	if err!=nil{
		t.Errorf("AddComments err   =%v",err)
	}
}

func testListComments(t *testing.T){
	vid:="123456"
	from:=1514764800
	to,err1:=strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000,10))
	log.Print(to)
	if err1!=nil{
		t.Errorf("testListComments error")
	}

	res,err:=listComments(vid,from,to)
	if err!=nil||res==nil{
		t.Errorf("Error of ListComments: %v",err)
		return
	}
	for i,ele:=range res{
		log.Printf("comment: %d,   %v \n",i,ele)
	}
}
