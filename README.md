# go_web_02
Heroku web app in golang.

## 1. golang dependency tool `dep`
```bash
$ brew install dep

$ mkdir $GOPATH/src/github.com/rwibawa/go_web_02
$ cd $GOPATH/src/github.com/rwibawa/go_web_02
$ dep init
$ dep status
$ dep ensure

$ vi Gopkg.toml
[metadata.heroku]
  root-package = "github.com/rwibawa/go_web_02"
  go-version = "go1.9"
```

Refer to [heroku go buildpack](https://github.com/heroku/heroku-buildpack-go)

## 2. web app
```bash
$ export PORT=8098
$ vi server.go
$ dep ensure
$ go build server.go
$ ./server
```

Open [http://localhost:8098/some/page/ryan](http://localhost:8098/some/page/ryan)

## 3. heroku
```bash
$ git init
$ git add -A
$ git commit -m "init repo"

$ heroku create go-web-02
$ git push heroku master

$ heroku open
$ heroku logs
```

Open [https://go-web-02.herokuapp.com/some/page/ryan](https://go-web-02.herokuapp.com/some/page/ryan)
