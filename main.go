package main

import (
	"log"
	"net/http"
	"flag"

	"github.com/NoCapCbas/webadmin/controllers"
	"github.com/NoCapCbas/webadmin/data"
)

func main() {
	dn := flag.String("driver", "postgres", "name of the database driver to use, postgres or mongo are supported")
	ds := flag.String("datasource", "", "database connection string")
  env := flag.String("env", "dev", "environment type, dev or prod")
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
	api.DB = db

  if *env == "dev" {
    log.Println(*env)
    // data.SeedDatabase(db.Connection)
  }

	if err := http.ListenAndServe(":8080", api); err != nil {
		log.Println(err)
	}
}
