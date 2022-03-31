package main

import (
	"log"
	"os"
	"time"

	_ "time/tzdata"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
	"github.com/szpp-dev-team/gakujo-api/model"
)

var (
	jUsername  string
	jPassword  string
	faculty    string
	department string
	course     string
	grade      string
	kamokuCode string
	classCode  string
	unit       string
	radio      string
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	jUsername = os.Getenv("J_USERNAME")
	jPassword = os.Getenv("J_PASSWORD")
	faculty = os.Getenv("FACULTY")
	department = os.Getenv("DEPARTMENT")
	course = os.Getenv("COURSE")
	grade = os.Getenv("GRADE")
	kamokuCode = os.Getenv("KAMOKU_CODE")
	classCode = os.Getenv("CLASS_CODE")
	unit = os.Getenv("UNIT")
	radio = os.Getenv("RADIO")
}

func main() {
	gc := gakujo.NewClient()
	if err := gc.Login(jUsername, jPassword); err != nil {
		log.Fatal(err)
	}
	kc, err := gc.NewKyoumuClient()
	if err != nil {
		log.Fatal(err)
	}
	formdata := &model.PostKamokuFormData{
		Faculty:    faculty,
		Department: department,
		Course:     course,
		Grade:      grade,
		KamokuCode: kamokuCode,
		ClassCode:  classCode,
		Unit:       unit,
		Radio:      radio,
	}

	c := gocron.NewScheduler(time.Local)
	_, err = c.Every(10 * time.Second).Do(func() {
		if 3 <= time.Now().Hour() && time.Now().Hour() <= 5 {
			return
		}
		if err := kc.PostRishuRegistration(formdata); err != nil {
			log.Println(err)
			return
		}
		log.Printf("%v を勝ち取りました！確認してください！\n", os.Getenv("KAMOKU_CODE"))
		c.Stop()
	})
	if err != nil {
		log.Fatal(err)
	}
	c.StartBlocking()
}
