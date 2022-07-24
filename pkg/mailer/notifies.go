package mailer

import (
	"encoding/json"
	"notifications/pkg/database"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func GetNotSentNotifies() []Notificationsdto {
	return fetchNotifies(`SELECT id, receipents, senddate, sentdate, issent, notificationtypeid FROM public.notifications WHERE issent=false;`)
}

func GetAllNotifies() []Notificationsdto {
	return fetchNotifies(`SELECT id, receipents, senddate, sentdate, issent, notificationtypeid FROM public.notifications;`)
}

func GetNotify(notifyid int) Notificationsdto {
	res := fetchNotifies("SELECT id, receipents, senddate, sentdate, issent, notificationtypeid FROM public.notifications WHERE id=" + strconv.Itoa(notifyid))
	return res[0]
}

func fetchNotifies(query string) []Notificationsdto {
	dtos := []Notificationsdto{}
	item := Notifications{}
	rows := fetchRows(query)
	for rows.Next() {
		err := rows.StructScan(&item)
		if err != nil {
			panic(err)
		}
		receipents := []Receipents{}
		json.Unmarshal(([]byte(item.Receipents)), &receipents)
		attachs := fetchAttachs(int(item.Id))
		body := fetchBody(int(item.Id))
		res := Notificationsdto{
			Id:                     item.Id,
			Receipents:             receipents,
			Senddate:               item.Senddate,
			Sentdate:               item.Sentdate,
			Issent:                 item.Issent,
			Notificationtypeid:     item.Notificationtypeid,
			Notificationbody:       body,
			Notificationattacments: attachs,
		}
		dtos = append(dtos, res)
	}
	return dtos
}

func fetchRows(query string) sqlx.Rows {
	db := database.Connect()

	rows, err := db.Queryx(query)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return *rows
}

func fetchBody(notifyid int) Notificationbody {
	rows := fetchRows(`SELECT id, notificationid, title, body FROM public.notificationbody WHERE notificationid=` + strconv.Itoa(notifyid))
	item := Notificationbody{}
	for rows.Next() {
		err := rows.StructScan(&item)
		if err != nil {
			panic(err)
		}
	}
	return item
}

func fetchAttachs(notifyid int) []Notificationattacments {
	db := database.Connect()
	rows, err := db.Queryx("SELECT id, notificationid, filename, summary, fileext, physicalpath, filecontent FROM public.notificationattacments WHERE notificationid=" + strconv.Itoa(notifyid))
	if err != nil {
		panic(err)
	}
	attach := Notificationattacments{}
	attachs := []Notificationattacments{}
	for rows.Next() {
		err := rows.StructScan(&attach)
		if err != nil {
			panic(err)
		}
		attachs = append(attachs, attach)
	}
	return attachs
}

func ExistNotify() int {
	res := GetNotSentNotifies()
	return len(res)
}

func MarkAsProcessed(notifyid int) error {
	db := database.Connect()
	_, err := db.Exec(`UPDATE public.notifications SET sentdate=NOW(), issent=$1 WHERE id=$2;`, true, notifyid)
	return err
}
