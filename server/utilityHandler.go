package server

import (
	"fmt"
	"net/http"
	"github.com/go-zoo/bone"
	"io/ioutil"
)

func defaultHandler(rw http.ResponseWriter, req *http.Request) {
	file, _ := ioutil.ReadFile("dist/index.html")
	rw.Write(file)
}

// /api/v1/ping
func getPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Site is up!")
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %v\n", bone.GetValue(r, "id"))
}