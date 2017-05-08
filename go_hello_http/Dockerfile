from golang:1.8.1-alpine

# add gcc and friends
RUN apk add --update alpine-sdk

# add bash
RUN apk add bash bash-doc bash-completion

# add go shim cli tools
ADD https://github.com/kelseyhightower/google-cloud-functions-go/releases/download/0.0.1/cloud-functions-go /bin
#ADD https://github.com/kelseyhightower/google-cloud-functions-go/releases/download/0.0.1/cloud-functions-go-shim /bin
ADD cloud-functions-go-shim /bin
RUN chmod 755 /bin/cloud-functions-*

# shell
CMD ["/bin/bash"]
