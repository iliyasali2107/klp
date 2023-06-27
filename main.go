package main

import (
	"klp/storage/pg_storage"
	"klp/worker"
	"log"
)

func main() {
	env, err := worker.InitEnv()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("env config is setted...")

	_, err = pg_storage.InitDB(env)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB is connected..")

	// mux := http.NewServeMux()

	// mux.HandleFunc("/ready")
}
