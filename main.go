package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/slack-go/slack"
	"github.com/szpp-dev-team/gakujo-api/gakujo"
	"github.com/szpp-dev-team/gakujo-api/model"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		if os.Getenv("SLACK_TOKEN") == "" {
			log.Fatalf(".env was not found\n%v\n", err)
		}
	}
}

func main() {
	c := cron.New()
	if _, err := c.AddFunc("00 0 * * *", task); err != nil {
		log.Fatal(err)
	}
	c.Start()
}


func task() {
	sc := slack.New(os.Getenv("SLACK_TOKEN"))
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
		Pretext: "人気な抽選科目(80%以上)を発表するよー！",
		Title:   "人気な抽選科目ランキング",
		Fields:  make([]slack.AttachmentField, 0),
	}

	sort.Slice(rows, func(i, j int) bool {
		percent1 := float64(rows[i].RegistrationStatus.FirstChoiceNum) / float64(rows[i].Capacity)
		percent2 := float64(rows[j].RegistrationStatus.FirstChoiceNum) / float64(rows[j].Capacity)
		return percent1 > percent2
	})

	for i, row := range rows {
		percent := float64(row.RegistrationStatus.FirstChoiceNum) / float64(row.Capacity) * 100
		if percent < 80 {
			break
		}
		attachment.Fields = append(attachment.Fields,
			slack.AttachmentField{
				Value: fmt.Sprintf(
					"%d位\t%s(%s)\t%.1f(%v/%v)%%",
					i+1,
					row.SubjectName,
					row.ClassName,
					percent,
					row.RegistrationStatus.FirstChoiceNum,
					row.Capacity,
				),
			},
		)
	}

	return &attachment
}
