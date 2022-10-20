package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	"github.com/labstack/echo/v4"
)

type AppointmentsResponse struct {
	Appointments []Appointment `json:"appointments"`
	Status       string        `json:"status"`
	Message      string        `json:"message"`
}

const (
	SuccessResponse      string = "success"
	ConflictResponse     string = "conflict"
	NotFoundResponse     string = "notFound"
	UnauthorizedResponse string = "unauthorized"
)

type Appointment struct {
	Item string
	Date time.Time
}

// var appoint2 []Appointment

// func booking() {
// 	t, _ := time.Parse("2006-01-02", "2022-10-19")
// 	appoint2 = append(appoint2, Appointment{Username: "Wilson", Date: t}, Appointment{Username: "Tony", Date: t})
// }

func createAppointments(c echo.Context) error {
	db, err := bolt.Open("DEMO.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var tx = &bolt.Tx{}
	b := tx.Bucket([]byte("Appointments"))
	username := c.FormValue("username")
	selectedItem := c.FormValue("selectedItem")
	selectedDate := c.FormValue("selectedDate")
	t, err := time.Parse("2006-01-02", selectedDate)
	if err != nil {
		return err
	}
	appoint2 := Appointment{}
	errMessage := fmt.Sprintf("%s，此日期已被預訂，請您重新選擇其他日期！", username)
	successMessage := fmt.Sprintf("預約成功！%s，您的預約日期為： %s", username, t.Format("2006-01-02"))

	b.ForEach(func(k, v []byte) error {
		json.Unmarshal(v, &appoint2)
		if appoint2.Date == t && appoint2.Item == selectedItem {
			return c.JSON(http.StatusConflict, AppointmentsResponse{Status: ConflictResponse, Message: errMessage})
		}
		return nil
	})
	data, _ := json.Marshal(appoint2)
	err = b.Put([]byte(username), data)
	return c.JSON(http.StatusOK, AppointmentsResponse{Status: SuccessResponse, Message: successMessage})
}

func searchAppointments(c echo.Context) error {

}

func main() {
	// Create a database named "DEMO.db" in your current directory.
	// It will be created if it doesn't exist.
	// And keep in connected.
	db, err := bolt.Open("DEMO.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Create a bucket(table) named "appointments".
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Appointments"))
		if err != nil {
			return fmt.Errorf("create bucket err: %s", err)
		}
		return nil
	})
	// Restore all appointments from "DEMO.db/Appointments"
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("Appointments"))
		// Loop all keys, values turn into (string), (time.Time) and append to appoint2
		b.ForEach(func(k, v []byte) error {
			t, err := time.Parse("2006-01-02", string(v))
			if err != nil {
				return err
			}
			appoint2 = append(appoint2, Appointment{Username: string(k), Date: t})
			return nil
		})
		fmt.Printf("從資料庫倒出來的預約記錄：%s", appoint2)
		return nil
	})
	// Close db connected.
	db.Close()

	// Book two sample dates.
	// booking()

	// Store all appointments to "DEMO.db/Appointments"
	for _, a := range appoint2 {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Appointments"))
			// b.Put([]byte(key),[]byte(value))
			err := b.Put([]byte(a.Username), []byte(a.Date.Format("2006-01-02")))
			return err
		})
	}
}
