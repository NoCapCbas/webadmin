package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/NoCapCbas/webadmin/controllers"
	"github.com/NoCapCbas/webadmin/data"
)

func main() {
	dn := flag.String("driver", "mongo", "name of the database driver to use, postgres or mongo are supported")
	ds := flag.String("datasource", "mongodb://localhost:27017", "database connection string")
	env := flag.String("env", "dev", "environment type, dev or prod")
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()

	if len(*dn) == 0 || len(*ds) == 0 {
		flag.Usage()
		return
	}
	api := controllers.NewAPI()

	// open the database connection
	db := &data.DB{}
	if err := db.Open(*dn, *ds); err != nil {
		log.Fatal("unable to connect to the database:", err)
	}
	log.Println("Database connection successful")
	api.DB = db

	if *env == "dev" {
		log.Printf("%s environment detected\n", *env)
		log.Println("Seeding development database")
		data.SeedDatabase(db.Connection)
	}

	log.Printf("Starting server on port %s\n", *port)
	if err := http.ListenAndServe(":"+*port, api); err != nil {
		log.Println(err)
	}
}
