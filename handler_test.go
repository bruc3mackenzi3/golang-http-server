
package main

import (
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
This module is for testing the Handlers for the HTTP server.

There is currently an issue with sending the post data.  multipart/form-data
isn't being properly set as it is with curl, throwing an error in the CSV
parser.

The code is being retained here to demonstrate testing web servers as a concept.
*/

var baseUrl string = "http://localhost:8080/"
var testMatrix string = "1,2,3\n4,5,6\n7,8,9"

func TestEchoHandle(t *testing.T) {
	t.Skip("HTTP Tests not working")
	req := httptest.NewRequest(
			"POST", baseUrl + "echo",
			strings.NewReader(testMatrix))
	req.Header.Add("Content-Type", "multipart/form-data")  // missing boundary
	w := httptest.NewRecorder()

	EchoHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
