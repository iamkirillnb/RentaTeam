package main

import (
	"flag"
	"fmt"
	"github.com/iamkirillnb/Rentateam/internal"
	"github.com/iamkirillnb/Rentateam/internal/handlers"
	"github.com/iamkirillnb/Rentateam/internal/repos"
	"github.com/iamkirillnb/Rentateam/pkg/logging"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


var cfgPath string

func init(){
	flag.StringVar(&cfgPath, "config", "dev.yaml", "config file path")
}

func main() {
	fmt.Println("Rentateam APP start")

	flag.Parse()

	log := logging.NewLogger()
	cfg := internal.GetConfig(cfgPath)

	connStr := cfg.DbConfig.DbUrl()

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("connection to DB fail")
	}

	// get blog repo
	repoLaw := repos.NewDbLaw(db, log)

	// start handler
	_, err = handlers.NewHandler(cfg, repoLaw, log)
	if err != nil {
		log.Fatal("make handler fail")
	}







}
