package main

import (
    "github.com/surrexi/learning-golang/api.golang/src/system/app"
    "flag"
    "github.com/joho/godotenv"
    "os"
    "github.com/surrexi/learning-golang/api.golang/src/system/database"
)

var port string
var dbHost string
var dbPort string
var dbUser string
var dbPass string
var dbOptions string
var dbDatabase string

func init() {
    flag.StringVar(&port, "port", "8000", "Assigning the port than the server should listen on.")
    flag.StringVar(&dbHost, "dbHost", "localhost", "Set the host for the application.")
    flag.StringVar(&dbPort, "dbPort", "3306", "Set the port for the application.")
    flag.StringVar(&dbUser, "dbUser", "test", "Set the user for the application.")
    flag.StringVar(&dbPass, "dbPass", "test", "Set the password for the application.")
    flag.StringVar(&dbOptions, "dbOptions", "parseTime=true", "Set options for the application.")
    flag.StringVar(&dbDatabase, "dbDatabase", "test", "Set the database for the application.")

    flag.Parse()

    if err := godotenv.Load("config.ini"); err != nil {
        panic(err)
    }

    if envPort := os.Getenv("PORT"); len(envPort) > 0 {
        port = envPort
    }

    if host := os.Getenv("PORT"); len(host) > 0 {
        dbHost = host
    }

    if port := os.Getenv("PORT"); len(port) > 0 {
        dbPort = port
    }

    if user := os.Getenv("PORT"); len(user) > 0 {
        dbUser = user
    }

    if pass := os.Getenv("PORT"); len(pass) > 0 {
        dbPass = pass
    }

    if options := os.Getenv("PORT"); len(options) > 0 {
        dbOptions = options
    }

    if db := os.Getenv("PORT"); len(db) > 0 {
        dbDatabase = db
    }
}

func main() {
    db, err := database.Connect(dbHost, dbPort, dbUser, dbPass, dbDatabase, dbOptions)
    if err != nil {
        panic(err)
    }

    s := app.NewServer()

    s.Init(port, db)
    s.Start()
}
