# 抽選絶対取る tool

抽選で落とした科目の履修登録を 10s 間隔で行います。

## Usage

1. `.env` を作成

```text
J_USERNAME=学情のid
J_PASSWORD=学情のpswd
KAMOKU_CODE=KAMOKU_CODE 等の確認方法を参照
CLASS_CODE=KAMOKU_CODE 等の確認方法を参照
UNIT=KAMOKU_CODE 等の確認方法を参照
RADIO=KAMOKU_CODE 等の確認方法を参照
YOUBI=KAMOKU_CODE 等の確認方法を参照
JIGEN=KAMOKU_CODE 等の確認方法を参照
```

2. 実行

```console
$ go run .
```

## `KAMOKU_CODE` 等の確認方法

1. 「一般講義履修登録」を「教務システム」から開きます

<img src="https://i.imgur.com/Qp3IOO4.png" width=400px>


2. 取りたい科目のコマの「鉛筆アイコン」をクリックします

<img src="https://i.imgur.com/rlmpgqn.png" width=400px>


3. 開発者ツールを「F12」キーで開きます

<img src="https://i.imgur.com/5PdzJe4.jpg" width=400px>

4. 取りたい科目をラジオボタンで選択して「登録」を行ってください。

5. リクエストのログが色々出てきますが、その内一番上の `searchKamoku.do` を選択し、「ペイロード(もしくは要求)」タブを選択してください。

<img src=https://i.imgur.com/mipP3aL.png width=400px>

6. ペイロードの中身を見ながら `.env` の対応する箇所に記入をしてください。(`J_USERNAME` や `J_PASSWORD` は学情のログイン時の情報です。)
