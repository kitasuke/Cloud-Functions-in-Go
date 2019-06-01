# Cloud-Functions-in-Go
sample projects

# Cloud Functions

## How to deploy Functions

```
$ gcloud functions deploy FUNC-NAME --runtime go111 --trigger-http
```

# Storage

## How to create Storage bucket

```
$ gsutil mb -c nearline gs://BUCKET-NAME
```

## How to deploy Functions for the Storage

```
$ gcloud functions deploy FUNC-NAME --runtime go111 \
    --trigger-resource BUCKET-NAME \
    --trigger-event google.storage.object.finalize
```

## How to upload files to the Storage

```
$ touch empty.txt && gsutil cp empty.txt gs://BUCKET-NAME
```

## How to show logs

```
$ gcloud beta functions logs read --limit 50
```

# Pub/Sub

## How to create Topic

```
$ gcloud pubsub topics create TOPIC-NAME
```

## How to deploy Functions for PubSub

```
$ gcloud functions deploy FUNC-NAME --runtime go111 --trigger-topic TOPIC-NAME
```

## How to publish topic to PubSub

```
$ gcloud pubsub topics publish TOPIC-NAME \
    --message '{"name": "foo"}'
```

## How to show logs

```
$ gcloud beta functions logs read --limit 50
```

# Storage

## How to deploy with env file

```
$ gcloud functions deploy TriggerHTTPBucket --runtime go111 \
    --trigger-http --env-vars-file .env.yaml
```

# BigQuery

## How to create table

```
$ bq mk TABLE-NAME
```

```
$ bq mk --table TABLE-NAME.TABLE-NAMe TABLE-NAME.json
```
