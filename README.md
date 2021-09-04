# chusen-kamoku-bot

第一志望の割合が80%以上の抽選科目のランキングを毎日 00:00 に slack に投げます。あんまり役立たない

## Deploy

### 0. slack-api で app を追加する

やる

### 1. `.env` を用意する

.env

```text
J_USERNAME=学情のusername
J_PASSWORD=学情のpassword
SLACK_CHANNEL_ID=slack のチャンネルid
SLACK_TOKEN=slack-bot の token
```

### 2. ビルド & 実行

golang 内で cron によるタスクスケジューリングはしてあるので、daemon 実行すれば定期実行は可能となります。

```console
$ go build -o ./app
$ ./app &
$ nohup ./app & # サーバー上でデプロイする場合
```
