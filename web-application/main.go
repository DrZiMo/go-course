package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	userRepo UserRepository
}

func main() {
	mux := http.NewServeMux()
	db, err := ConnectToDB("api.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		infoLog:  log.New(os.Stderr, "INFO\t", log.Ltime|log.LstdFlags),
		userRepo: NewSQLRepository(db),
	}

	// mux.HandleFunc("/", home)
	// mux.HandleFunc("/about", about)
	// mux.HandleFunc("/contact", contact)

	fmt.Println("Listening to port 8080 ...")
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}

}
