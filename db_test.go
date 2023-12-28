package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const DotenvPath = ".env"

func TestDatabaseInit(t *testing.T) {
    err := godotenv.Load(DotenvPath); if err != nil { t.Fatalf(err.Error()) }
    dbName := os.Getenv("DB_NAME")
    
    dbClient, err := InitDatabase(DotenvPath)
    if err != nil { t.Fatalf("Unable to create dbclient: %v", err) }
    defer dbClient.Close()
    rows, err := dbClient.Query("SHOW DATABASES LIKE " + fmt.Sprintf(`"%s"`,dbName))
    if err != nil { t.Fatalf("Database check exsistance failed: %v", err) }
    defer rows.Close()

    if !rows.Next() { 
        t.Fatalf("Database doesnt exist after creation: %v", err) 
    } else {
        _, err := dbClient.Exec("DROP DATABASE IF EXISTS " + dbName)
        if err != nil { t.Fatalf("Unable to delete database: %v", err) }
        t.Logf("Database (%s) Created and Deleted seccussfully!", dbName)
    }
}

