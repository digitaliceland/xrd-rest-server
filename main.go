package main

import (
	"flag"
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"os"
	"time"
)


type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

// Middleware
func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

// Handlers
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "pong"}`))
}

func (app *application) time(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	t := time.Now().UTC()
	timestring := t.Format(time.RFC3339)
	msg := fmt.Sprintf(`{"time": "%s"}`, timestring)
	w.Write([]byte(msg))
}

//Routes
func (app *application) routes() http.Handler {
	standardMiddleware := alice.New(app.logRequest)

	mux := pat.New()
	mux.Get("/", standardMiddleware.ThenFunc(app.home))
	mux.Get("/ping", standardMiddleware.ThenFunc(app.ping))
	mux.Get("/time", standardMiddleware.ThenFunc(app.time))

	return standardMiddleware.Then(mux)
}


func main() {
	addr := flag.String("addr", ":1234", "HTTP network address to bind to")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
