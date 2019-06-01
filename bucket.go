package Cloud_Functions_in_Go

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

func TriggerHTTPBucket(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		names, err := r.URL.Query()["name"]
		if !err || len(names[0]) < 1 {
			fmt.Fprintf(w, "no name found in param")
			return
		}
		WriteBucket(w, names[0])
	default:
		http.Error(w, "405 - Method not allowd", http.StatusMethodNotAllowed)
	}
}

func WriteBucket(w http.ResponseWriter, name string) {

	bucketName := os.Getenv("BUCKET_NAME")

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Fprintf(w, "Storage error: %T %s", err, err)
		return
	}

	defer client.Close()

	objectName := time.Now().Format("20060102150405")

	fw := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
	if _, err := fw.Write([]byte(name + "\r\n")); err != nil {
		fmt.Fprintf(w, "Write error %T %s", err, err)
		return
	}

	if err := fw.Close(); err != nil {
		fmt.Fprintf(w, "Close error %T %s", err, err)
		return
	}
}
