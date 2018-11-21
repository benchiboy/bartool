// ProcFlow.go
package control

import (
	"bartool/comm"
	"bartool/sessions"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func LoginShow(w http.ResponseWriter, req *http.Request) {
	buf, _ := ioutil.ReadFile("./htmls/login.html")
	w.Write(buf)
}

/*
	登录
*/
func Login(w http.ResponseWriter, req *http.Request) {
	log.Println("Login......")
	var login comm.Login_Node
	err := json.NewDecoder(req.Body).Decode(&login)
	if err != nil {
		log.Println("Login......", err.Error())
		return
	}
	defer req.Body.Close()
	log.Println(login)
	if login.User_Pwd != "123456" {
		http.Error(w, fmt.Errorf("用户或密码错误！").Error(), http.StatusInternalServerError)
		return
	}
	log.Println("login ok ... ")
	sessions.SetLogined(w, req)
	return

}

/*
	退出
*/
func Signout(w http.ResponseWriter, req *http.Request) {
	log.Println("Signout......")
	sessions.LogoutFunc(w, req)
}

/*
	修改密码
*/
func ChangePwd(w http.ResponseWriter, req *http.Request) {
	var change comm.ChangePwd_Node
	err := json.NewDecoder(req.Body).Decode(&change)
	if err != nil {
		log.Println("ResetPwd......", err.Error())
		return
	}
	defer req.Body.Close()
}
