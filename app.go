package main

import (
  "fmt"
  "github.com/JackC/pgx"
  "net/http"
  "os"
  "log"
  "encoding/json"
)

var pool *pgx.ConnectionPool

type Response map[string]interface{}

func (r Response) String() (string) {
  b, err := json.Marshal(r)
  if err != nil {
    return ""
  }
  return string(b)
}

func main() {
  var err error
  connectionOptions := pgx.ConnectionParameters{
    Host:     "localhost",
    User:     "piotrek",
    Password: "pass",
    Database: "rowlf_development"}
  poolOptions := pgx.ConnectionPoolOptions{MaxConnections: 10}
  pool, err = pgx.NewConnectionPool(connectionOptions, poolOptions)

  if err != nil {
    fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
    os.Exit(1)
  }
  http.HandleFunc("/", handler);
  err = http.ListenAndServe("localhost:3000", nil)
  if err != nil {
    log.Fatal(err)
  }
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  if lodgings, err := pool.SelectRows("select name, external_id FROM lodgings_lodgings LIMIT 100"); err == nil {
    fmt.Fprint(w, Response{"lodgings": lodgings})
  } else {
    http.Error(w, fmt.Sprintf("Internal server error: %s", err), http.StatusInternalServerError)
  }
}
