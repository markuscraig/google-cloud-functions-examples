// Copyright 2017 Google Inc. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/kelseyhightower/google-cloud-functions-go/event"
)

var (
	entryPoint string
	eventType  string
	pluginPath string
)

func main() {
	flag.StringVar(&entryPoint, "entry-point", "F", "the name of a Go function that will be executed when the Cloud Function is triggered.")
	flag.StringVar(&eventType, "event-type", "", "The Cloud Function event type. (bucket, http, or topic)")
	flag.StringVar(&pluginPath, "plugin-path", "functions.so", "The path to the Go plugin that exports the function to be executed.")
	flag.Parse()

	if eventType == "" {
		log.Fatal("Event type required; set using the --event-type flag.")
	}

	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("unable to load the event: %s", err)
	}

	var message string

	switch eventType {
	case "bucket":
		message, err = objectChangeHandler(stdin)
	case "http":
		message, err = httpHandler(stdin)
	case "topic":
		message, err = topicPublishHandler(stdin)
	default:
		log.Fatalf("invalid event type: %s", eventType)
	}

	if err != nil {
		log.Fatal(err)
	}

	if message != "" {
		os.Stdout.WriteString(message)
	}

	os.Exit(0)
}

func objectChangeHandler(data []byte) (string, error) {
	var e event.ObjectChange
	var message string

	err := json.Unmarshal(data, &e)
	if err != nil {
		return "", fmt.Errorf("unable to load the event: %s", err)
	}

	/*
		message, err = f.(func(event.ObjectChange) (string, error))(e)
		if err != nil {
			return "", err
		}
	*/
	message = "FIXME"

	return message, nil
}

func topicPublishHandler(data []byte) (string, error) {
	var e event.TopicPublish
	var message string

	err := json.Unmarshal(data, &e)
	if err != nil {
		return "", fmt.Errorf("unable to load the event: %s", err)
	}

	/*
		message, err = f.(func(event.TopicPublish) (string, error))(e)
		if err != nil {
			return "", err
		}
	*/
	message = "FIXME"

	return message, nil
}

func httpHandler(data []byte) (string, error) {
	var httpRequest event.HTTP

	err := json.Unmarshal(data, &httpRequest)
	if err != nil {
		return "", fmt.Errorf("unable to load the event: %s", err)
	}

	r := httptest.NewRequest(httpRequest.Method, httpRequest.URL, bytes.NewBufferString(httpRequest.Body))
	for k, v := range httpRequest.Header {
		r.Header.Add(k, v)
	}

	r.RemoteAddr = httpRequest.RemoteAddr

	w := httptest.NewRecorder()

	//f.(func(http.ResponseWriter, *http.Request))(w, r)
	myHandler(w, r)

	resp := w.Result()

	header := make(map[string]string)
	for k, v := range resp.Header {
		header[k] = strings.Join(v, ",")
	}

	out, err := json.Marshal(&event.HTTPResponse{
		Body:       w.Body.String(),
		Header:     header,
		StatusCode: resp.StatusCode,
	})

	if err != nil {
		return "", err
	}

	return string(out), nil
}

type response struct {
	Message string `json:"message"`
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	if err := json.NewEncoder(w).Encode(response{Message: string(d)}); err != nil {
		w.WriteHeader(500)
		return
	}
}
