// rcs_contract_mgr project main.go
package main

import (
	"context"

	"bartool/control"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	"syscall"

	"time"

	"bartool/sessions"

	gcontext "github.com/gorilla/context"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	http_srv *http.Server
)

func go_WebServer() {

	log.Println("Listen Service start...")
	http.Handle("/static/", http.FileServer(http.Dir("public")))

	http.HandleFunc("/index/", sessions.MustLogin(control.Index))
	http.HandleFunc("/logout/", control.Signout)
	http.HandleFunc("/login/", control.Login)
	http.HandleFunc("/gologin/", control.LoginShow)

	http.HandleFunc("/bardatas/", sessions.MustLogin(control.BarLoad))
	http.HandleFunc("/baradd/", sessions.MustLogin(control.Baradd))
	http.HandleFunc("/bardel/", sessions.MustLogin(control.Bardel))
	http.HandleFunc("/barget/", sessions.MustLogin(control.Barget))
	http.HandleFunc("/barset/", sessions.MustLogin(control.Barset))
	http.HandleFunc("/barpage/", sessions.MustLogin(control.BarLoad))
	http.HandleFunc("/barfind/", sessions.MustLogin(control.BarLoad))

	http_srv = &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: gcontext.ClearHandler(http.DefaultServeMux),
	}

	if err := http_srv.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}

}

func init() {
	log.Println("System Paras Init......")
	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
	log.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "rcs_contract.log",
		MaxSize:    500, // megabytes
		MaxBackups: 50,
		MaxAge:     90, //days
	}))

}

/*
  see doc
*/

func checkDB() {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	log.Println("Recv a kill signal and exit ...")

	time.Sleep(time.Second * 10)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := http_srv.Shutdown(ctx)
	log.Println("Server gracefully stopped:", err)

}

func main() {
	go_WebServer()
}
