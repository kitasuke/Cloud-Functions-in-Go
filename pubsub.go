package Cloud_Functions_in_Go

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
)

type PubSubMessage struct {
	Data []byte `json:"data"`
}

type Info struct {
	Name string `json:"name"`
}

func LogPubSubMessage(ctx context.Context, m PubSubMessage) error {
	var i Info

	err := json.Unmarshal(m.Data, &i)

	if err != nil {
		log.Printf("Error: %T message: %v", err, err)
		return nil
	}

	log.Printf("Message to %s", i.Name)
	return nil
}

func TriggerPubSubToBigquery(ctx context.Context, m PubSubMessage) error {
	var i Info

	err := json.Unmarshal(m.Data, &i)
	if err != nil {
		log.Printf("Error %T %v", err, err)
		return nil
	}

	InsertToBigQuery(ctx, &i)
	return nil
}

func InsertToBigQuery(ctx context.Context, i *Info) {
	projectID := os.Getenv("GCP_PROJECT")

	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("BigQuery connection error %T %v", err, err)
		return
	}

	defer client.Close()

	u := client.Dataset("GREETINGS").Table("NAMES").Inserter()

	items := []*Info{i}
	err = u.Put(ctx, items)
	if err != nil {
		log.Printf("Data write error %T %v", err, err)
		return
	}
}
