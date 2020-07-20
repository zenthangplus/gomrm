package main

import (
    "fmt"
    "github.com/zenthangplus/goorm"
)

func main() {
    db, err := gomrm.Connect("mysql", "root:secret@tcp(localhost:3306)/gomrm_test")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    fmt.Println("EXAMPLE FOR COLLECTION STYLE")
    collection, err := db.Query("SELECT * from customers")
    if err != nil {
        panic(err)
    }

    for collection.Next() {
        result := collection.Get()
        fmt.Printf("ID [%s], Name [%s]\n", result["id"], result["name"])
    }

    firstResult := collection.First()
    fmt.Printf("First ID [%s], Name [%s]\n", firstResult["id"], firstResult["name"])

    lastResult := collection.Last()
    fmt.Printf("Last ID [%s], Name [%s]\n", lastResult["id"], lastResult["name"])

    fmt.Println("EXAMPLE FOR MAP STYLE")
    results, err := db.QueryRaw("SELECT * from customers")
    if err != nil {
        panic(err)
    }
    for _, result := range results {
        fmt.Printf("ID [%s], Name [%s]\n", result["id"], result["name"])
    }
}
