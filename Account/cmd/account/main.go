package main

import (
	"github.com/kckurkcu1/GO-Microservices-Project/Account"
	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
	"log"
	"time"
)

type Config struct {
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r Account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = Account.NewPostgresRepository(cfg.DatabaseURL)
		if err != nil {
			log.Println(err)
		}
		return
	})
	defer r.Close()
	log.Println("Listening on port 8080...")
	s := Account.NewService(r)
	log.Fatal(Account.ListenGRPC(s, 8080))
}
