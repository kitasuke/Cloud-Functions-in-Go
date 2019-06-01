package Cloud_Functions_in_Go

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/functions/metadata"
)

type GCSEEvent struct {
	Bucket         string    `json:"bucket"`
	Name           string    `json:"name"`
	Metageneration string    `json:"metageneration"`
	ResourceState  string    `json:"resourceState"`
	TimeCreated    time.Time `json:"timeCreated"`
	Updated        time.Time `json:"updated"`
}

func LogStorageEvent(ctx context.Context, e GCSEEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}

	log.Printf("%v", meta.EventID)
	log.Printf("%v", meta.EventType)
	log.Printf("%v", e.Bucket)
	log.Printf("%v", e.Name)
	log.Printf("%v", e.ResourceState)
	log.Printf("%v", e.TimeCreated)
	log.Printf("%v", e.Updated)
	return nil
}
