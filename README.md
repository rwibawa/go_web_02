# go_web_02
Heroku web app in golang.

## 1. golang dependency tool `dep`
```bash
$ brew install dep

$ cd $GOPATH/src/github.com/rwibawa/go_web_01
$ dep init
$ dep status
$ dep ensure
```

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


