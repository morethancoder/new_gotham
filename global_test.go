package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)


func TestGlobalValues(t *testing.T) { 
    err := godotenv.Load(DotenvPath); if err != nil { t.Fatalf(err.Error()) }
    dbName := os.Getenv("DB_NAME")
    
    dbClient, err := InitDatabase(DotenvPath); if err != nil {t.Fatalf(err.Error())}
    err = InitGlobalValuesTable(dbClient); if err != nil {t.Fatalf(err.Error())}
    i := GlobalValuesInstance{
        Count: 0,
    }
    err = i.Create(dbClient); if err != nil {t.Fatalf(err.Error())}
    var foo GlobalValuesInstance = GlobalValuesInstance{ID: 1}

    err = foo.Create(dbClient); if err != nil {t.Fatal(err.Error())} 
    err = foo.Read(dbClient); if err !=  nil {t.Fatalf(err.Error())}
    t.Logf("GlobalCount: ID:%d Count:%d", i.ID, i.Count)
    err = foo.Update(dbClient); if err != nil {t.Fatalf(err.Error())}
    t.Logf("GlobalCount: ID:%d Count:%d", foo.ID, foo.Count)
    if foo.Count == 0 {
        t.Fatalf("Incremented count shouldnt equal to zero")
    }
    err = foo.Delete(dbClient); if err != nil {t.Fatalf(err.Error())}
    _, err = dbClient.Exec("DROP TABLE "+ GlobalValuesTableKey); if err != nil {t.Fatalf(err.Error())}
    _, err = dbClient.Exec("DROP DATABASE IF EXISTS "+ dbName); if err != nil {t.Fatalf(err.Error())}

}
