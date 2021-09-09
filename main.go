package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
	"github.com/szpp-dev-team/gakujo-api/model"
)

var (
	kamokuCode string
	classCode  string
	unit       int
	radio      int
	youbi      int
	jigen      int
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		if os.Getenv("SLACK_TOKEN") == "" {
			log.Fatalf(".env was not found\n%v\n", err)
		}
	}
	kamokuCode = os.Getenv("KAMOKU_CODE")
	classCode = os.Getenv("CLASS_CODE")
	unit, _ = strconv.Atoi(os.Getenv("UNIT"))
	radio, _ = strconv.Atoi(os.Getenv("RADIO"))
	youbi, _ = strconv.Atoi(os.Getenv("YOUBI"))
	jigen, _ = strconv.Atoi(os.Getenv("JIGEN"))
}

func main() {
	c := cron.New()
	if _, err := c.AddFunc("15 * * * *", task); err != nil {
		log.Fatal(err)
	}
	c.Start()

	for {
		time.Sleep(time.Hour * 24)
		log.Println("1 day later...")
	}
}

func task() {
	log.Println(os.Getenv("J_USERNAME"))
	log.Println(os.Getenv("J_PASSWORD"))
	gc := gakujo.NewClient()
	if err := gc.Login(os.Getenv("J_USERNAME"), os.Getenv("J_PASSWORD")); err != nil {
		log.Fatal(err)
	}
	kc, err := gc.NewKyoumuClient()
	if err != nil {
		log.Fatal(err)
	}
	formdata := model.NewPostKamokuFormData(kamokuCode, classCode, unit, radio, youbi, jigen)
	if err := kc.PostRishuRegistration(formdata); err != nil {
		if errors.Is(err, gakujo.OverCapasityError{}) {
			log.Println("定員オーバーで履修登録ができませんでした。諦めない")
			return
		}
		log.Println("別のエラーが発生したようです")
		log.Println(err)
		return
	}
	log.Printf("%v を勝ち取りました！確認してください！\n", os.Getenv("KAMOKU_CODE"))
	os.Exit(0)
}
