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

## 4. angular4
```bash
$ ng new webui
  create webui/README.md (1021 bytes)
  create webui/.angular-cli.json (1240 bytes)
  create webui/.editorconfig (245 bytes)
  create webui/.gitignore (516 bytes)
  create webui/src/assets/.gitkeep (0 bytes)
  create webui/src/environments/environment.prod.ts (51 bytes)
  create webui/src/environments/environment.ts (387 bytes)
  create webui/src/favicon.ico (5430 bytes)
  create webui/src/index.html (292 bytes)
  create webui/src/main.ts (370 bytes)
  create webui/src/polyfills.ts (2667 bytes)
  create webui/src/styles.css (80 bytes)
  create webui/src/test.ts (1085 bytes)
  create webui/src/tsconfig.app.json (211 bytes)
  create webui/src/tsconfig.spec.json (304 bytes)
  create webui/src/typings.d.ts (104 bytes)
  create webui/e2e/app.e2e-spec.ts (287 bytes)
  create webui/e2e/app.po.ts (208 bytes)
  create webui/e2e/tsconfig.e2e.json (235 bytes)
  create webui/karma.conf.js (923 bytes)
  create webui/package.json (1310 bytes)
  create webui/protractor.conf.js (722 bytes)
  create webui/tsconfig.json (363 bytes)
  create webui/tslint.json (2985 bytes)
  create webui/src/app/app.module.ts (314 bytes)
  create webui/src/app/app.component.css (0 bytes)
  create webui/src/app/app.component.html (1120 bytes)
  create webui/src/app/app.component.spec.ts (986 bytes)
  create webui/src/app/app.component.ts (207 bytes)
Installing packages for tooling via npm.
Installed packages for tooling via npm.
Directory is already under version control. Skipping initialization of git.
Project 'webui' successfully created.

$ cd webui/
$ vi .angular-cli.json -> OutDirs: "../dist"
$ ng build
$ cd ../
$ go run go_web_02
```
