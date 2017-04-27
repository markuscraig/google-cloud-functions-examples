## Run the Docker Image

I put together a Docker image containing Go 1.8.1 a Go shim for Google Cloud Functions (by Kelsey Hightower of Google).

* https://hub.docker.com/r/barl3yb3ar/go1.8.1-gcf/

This command maps your local GOPATH to the docker image's "/go" directory...

```bash
$ docker run -v $GOPATH:/go -it --entrypoint bash barl3yb3ar/go1.8.1-gcf
```

## Build the Go Function as Zip File

```bash
docker$ cd $GOPATH/src/github.com/{YOUR-PROJECT-PATH}
docker$ go build -buildmode=plugin -o functions.so main.go
docker$ exit
```

## Deploy the Function Zip File

I had to manually create the function and upload the zip file from the web console.

There does not appear to be a way to do this using the 'gcloud' cli tool.

## Invoke the Function

Hmmm, this is not working yet...

```bash
$ curl -X POST https://us-central1-fithub-163017.cloudfunctions.net/go_hello_http -H "Content-Type: text/plain" --data 'Go Serverless!'
```
