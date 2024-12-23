package main

import (
	"github.com/kckurkcu1/GO-Microservices-Project/Order"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"log"
	"time"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
	AccountURL  string `envconfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL  string `envconfig:"CATALOG_SERVICE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	var r Order.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = Order.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println("Listening on port 8080 ...")
	s := Order.NewService(r)
	log.Fatal(Order.ListenGRPC(s, cfg.AccountURL, cfg.CatalogURL, 8080))
}
