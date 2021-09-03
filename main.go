package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
	"github.com/szpp-dev-team/gakujo-api/model"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	sc := slack.New(os.Getenv("SLACK_TOLEN"))
	gc := gakujo.NewClient()
	if err := gc.Login(os.Getenv("J_USERNAME"), os.Getenv("J_PASSWORD")); err != nil {
		log.Fatal(err)
	}

	kc, err := gc.NewKyoumuClient()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := kc.ChusenRegistrationRows()
	if err != nil {
		log.Fatal(err)
	}

	attachment := makeMessageAttachment(rows)
	_, _, err = sc.PostMessage(
		os.Getenv("SLACK_CHANNEL_ID"),
		slack.MsgOptionAttachments(*attachment),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func makeMessageAttachment(rows []*model.ChusenRegistrationRow) *slack.Attachment {
	attachment := slack.Attachment{
		Color:   "good",
		Pretext: "人気な抽選科目(75%以上)を発表するよー！",
		Title:   "人気な抽選科目ランキング",
		Fields:  make([]slack.AttachmentField, 0),
	}

	return &attachment
}
