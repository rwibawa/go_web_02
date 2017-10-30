package server

import (
	"os"
	"log"
	"net/http"

	//_ "github.com/lib/pq"
	//"github.com/jmoiron/sqlx"
	"github.com/go-zoo/bone"
	"github.com/urfave/negroni"
	//"github.com/jasonlvhit/gocron"

	//swagger "github.build.ge.com/Aviation/water-wash-optimizer-diagnostics-ww-events/swagger-ui"
)

const base_path_1 = "/api/v1/"

//var db *sqlx.DB
var mux = bone.New()
var port string
var isRunningLocally = false

func init() {
	//var err error

	// This line only works on Cloud Foundry.
	// It will return an empty string when run locally.
	port = os.Getenv("PORT")

	if port == "" {
		isRunningLocally = true

		// Connect to Postgres on Local Machine
		//db, err = sqlx.Open("postgres", "user=postgres password=password dbname=postgres sslmode=disable")
	}
}

// Define our Routes and their respective Handlers.
//
// Note the use of the "base_path_1" constant.
// It is the standard to have base paths to separate newer and older versions of our handlers,
// so old users can still use the older versions, while newer users can use the newer versions.
//
// With this in place, our URLs will now look like:
// localhost:3000/api/v1/myRoute
func SetupServer() {
	mux.Get("/some/page/:id", http.HandlerFunc(MyHandler))

	// Public Methods
	mux.GetFunc(base_path_1 + "ping", getPing)

	// Default page is "/dist/index.html"
	mux.Get("/*", http.FileServer(http.Dir("dist")))
}

func StartServer() {
	// If we're running the app locally, manually assign our own Port number.
	if isRunningLocally {
		port = "3000"
	}

	log.Print("Listening on Port " + port + "...")

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":"+port)
}