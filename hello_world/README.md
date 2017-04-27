## Get Active Project ID

```bash
$ gcloud info | grep Project
```

## Make a Bucket for Functions

```bash
$ gsutil mb -p <PROJECT-ID> 'gs://GLOBALLY-UNIQUE-BUCKET-NAME'
```

## Deploy Function

```bash
$ gcloud beta functions deploy helloWorld --stage-bucket <GLOBALLY-UNIQUE-BUCKET-NAME> --trigger-topic hello_world
```

## Invoke Function

```bash
$ gcloud beta functions call helloWorld --data '{"message":"Hello World!"}'
```

## Function Console Logs

```bash
$ gcloud beta functions logs read helloWorld
```

## Delete Function

```bash
$ gcloud beta functions delete helloWorld
```

## Links

* https://cloud.google.com/sdk/gcloud/reference/beta/functions/deploy
* https://cloud.google.com/functions/docs/calling/http
