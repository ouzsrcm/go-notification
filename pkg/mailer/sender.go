package mailer

import (
	"fmt"
	"net/smtp"
	"strconv"
)

var emailAuth smtp.Auth

func Run() bool {
	fmt.Println("started...")
	notifies := GetNotSentNotifies()
	fmt.Printf("fetched %s notify", strconv.Itoa(len(notifies)))
	for _, item := range notifies {
		item.send()
	}
	return true
}

func (n Notificationsdto) send() bool {
	res := false
	if n.Notificationtypeid == 1 {
		res = n.sendsms()
	} else {
		res = n.sendemail()
	}
	if res {
		n.markasprocessed()
		return true
	}
	return false
}

func (n Notificationsdto) sendemail() bool {
	fmt.Printf("%s mail sent.", n.Senddate)
	receipents := []string{}
	for _, item := range n.Receipents {
		receipents = append(receipents, item.Email)
	}
	sendemail(n.Notificationbody.Title, n.Notificationbody.Body, receipents)
	return true
}

func (n Notificationsdto) sendsms() bool {
	fmt.Printf("%s sms sent.", n.Senddate)
	receipents := []string{}
	for _, item := range n.Receipents {
		receipents = append(receipents, item.Email)
	}
	sendsms(n.Notificationbody.Title, n.Notificationbody.Body, receipents)
	return true
}

func (n Notificationsdto) markasprocessed() (bool, error) {
	err := MarkAsProcessed(int(n.Id))
	if err != nil {
		return false, err
	}
	return true, err
}

func sendemail(subject string, body string, receipents []string) (bool, error) {
	host := "smtp.gmail.com"
	from := "oguzhan.saricam@gmail.com"
	pwd := "1"
	port := "587"

	emailAuth = smtp.PlainAuth("", from, pwd, host)
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subj := "Subject: " + subject + "\n"
	msg := []byte(subj + mime + "\n" + body)
	addr := fmt.Sprintf("%s:%s", host, port)
	err := smtp.SendMail(addr, emailAuth, from, receipents, msg)
	if err != nil {
		return false, err
	}
	return true, nil
}

func sendsms(subject string, body string, receipents []string) (bool, error) {
	//TODO: ongoing...
	return true, nil
}
