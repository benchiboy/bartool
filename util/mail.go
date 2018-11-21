// util
package util

import (
	"crypto/tls"

	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/gomail.v2"
)

type Mail_Conf struct {
	User         string
	Passwd       string
	From_addr    string
	Smtp_addr    string
	Smtp_port    int
	Mail_Subject string
	Mail_From    string
	Mail_to      []string
	Mail_body    string
}

/*
	Mail constructor
*/
func New_Mail(smtp_addr string, smtp_port int, from string, user string, pwd string) *Mail_Conf {
	return &Mail_Conf{Smtp_addr: smtp_addr, Smtp_port: smtp_port, User: user, Passwd: pwd, From_addr: from}
}

/*
	Mail  send
*/

func (y *Mail_Conf) GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func (y *Mail_Conf) Send_Mail(tos []string, subject string, body string, embed string, attachs []string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", y.From_addr)
	m.SetHeader("To", tos...)
	m.SetHeader("Subject", subject)
	if embed != "" {
		m.Embed(embed)
	}
	for _, v := range attachs {
		m.Attach(v)
	}
	m.SetBody("text/html", body)
	d := gomail.NewDialer(y.Smtp_addr, y.Smtp_port, y.User, y.Passwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Println("send mail error:", err.Error())
		return err
	}
	log.Println("send mail succ!")
	return nil
}
