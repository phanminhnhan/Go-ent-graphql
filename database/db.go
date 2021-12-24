package database

import (
	"context"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"go-ent-graphql/ent"
)
var EntClient *ent.Client
func Init() {
	Client, err := ent.Open("postgres","host=localhost port=5432 user=root dbname=postgres password=1999 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database successfully")
	//return Client
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	EntClient = Client
}