package utils

import (
	"context"
	"log"

	"entgo.io/ent/examples/privacytenant/ent"
)

func Database(c context.Context) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(c); err != nil {
		log.Fatal(err.Error())
	}

	return client, err
}
