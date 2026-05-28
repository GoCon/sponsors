# gocon/sponsors

## 使い方

### Go1.26.3のインストール

```sh
$ go install golang.org/dl/go1.26.3@latest
$ go1.26.3 download
```

### 抽選

抽選会では本番においても`testdata/practice.csv`を使ってデモンストレーションを行う。

```sh
$ cat testdata/practice.csv | go1.26.3 run ./cmd/lottery | tee practice-result.txt
```

抽選会までに`applicants.csv`に参加企業とスポンサープランの一覧を出力しておく。
CSVの内容は以下のように、社名、プラン、外れた場合に次のプランの抽選に参加するかが含まれています。

```
company,plan,next
株式会社Gopher1,free,FALSE
株式会社Gopher2,platinum,TRUE
株式会社Gopher3,gold,TRUE
```

```sh
$ cat applicants.csv | go1.26.3 run ./cmd/lottery | tee result.txt
```
