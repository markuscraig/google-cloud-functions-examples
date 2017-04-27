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
$ gcloud beta functions deploy helloHttp --stage-bucket <GLOBALLY-UNIQUE-BUCKET-NAME> --trigger-http
```

## Get Function URL

```bash
$ gcloud beta functions describe helloHttp | grep url
```

## Invoke Function

```bash
$ curl -X POST \
    https://YOUR-URL?foo=bar1 \
    -H "Content-Type: application/json" \
    -H "X-MyHeader: hi" \
    --data '{"foo":"bar2","name":"Mark"}'
```

## Function Console Logs

```bash
$ gcloud beta functions logs read helloHttp
```

## Delete Function

```bash
$ gcloud beta functions delete helloHttp
```

## Links

* https://cloud.google.com/sdk/gcloud/reference/beta/functions/deploy
* https://cloud.google.com/functions/docs/calling/http
