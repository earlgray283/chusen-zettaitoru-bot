# 抽選絶対取る bot

〜For 抽選漏れした哀れな静大生〜  

15分毎に履修登録 api を叩きます

## Setup

`.env` を作ってください

```text
J_USERNAME=学情のid
J_PASSWORD=学情のpswd
KAMOKU_CODE=とりたい科目のコード
CLASS_CODE=とりたい科目のクラスコード(html から確認してください)
UNIT=とりたい科目の単位数
RADIO=とりたい科目が上から何番目にあるか(0-indexedで)
YOUBI=曜日(月曜日=1)
JIGEN=何コマ目
```

## Run

### Linux, macOS

バックグラウンドで動きます

```console
$ go build -o ./app
$ ./app &
```

### Windows

絶対に端末画面は消さないでください

```console
$ go run .
```
