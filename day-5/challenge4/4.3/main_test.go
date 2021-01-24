package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func init() {
	router = SetupRouter()
}

func TestCompleteApi(t *testing.T) {
	os.Setenv("REPOSITORY_TYPE", "Memory")

	doAllAPIRequests(t)
}

func doAllAPIRequests(t *testing.T) {
	doGetItems(router, t, "", true, 0)
	doGetItem(router, t, 0, "", false)

	doPostItem(router, t, "POST", `{"id":1,"title":"Test_1","isdone":true}`)
	doGetItems(router, t, "Test_1", true, 1)
	doGetItem(router, t, 1, "Test_1", true)

	doPostItem(router, t, "PUT", `{"id":2,"title":"Test_2","isdone":true}`)
	doGetItems(router, t, "Test_1", true, 2)
	doGetItem(router, t, 2, "Test_2", true)

	doDeleteItem(router, t, 2)
	doGetItems(router, t, "Test_1", true, 1)

	doPatchItem(router, t, 1, `{"id":1,"title":"Test_3","isdone":false}`)
	doGetItems(router, t, "Test_3", false, 1)
	doGetItem(router, t, 1, "Test_3", false)
}

func doPostItem(r http.Handler, t *testing.T, method string, payload string) {
	request := doRequest(r, method, "/api/", payload)

	assert.Equal(t, http.StatusCreated, request.Code)

	var response map[string]string

	err := json.Unmarshal([]byte(request.Body.String()), &response)
	value, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "OK", value)
}

func doGetItem(r http.Handler, t *testing.T, id int, title string, isdone bool) {
	request := doRequest(r, "GET", fmt.Sprintf("/api/%v", id), "")

	assert.Equal(t, http.StatusOK, request.Code)

	var response Item

	err := json.Unmarshal([]byte(request.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, response.ID, id)
	assert.Equal(t, response.Title, title)
	assert.Equal(t, response.IsDone, isdone)
}

func doGetItems(r http.Handler, t *testing.T, title string, isdone bool, length int) {
	request := doRequest(r, "GET", "/api/", "")

	assert.Equal(t, http.StatusOK, request.Code)

	var response []Item

	err := json.Unmarshal([]byte(request.Body.String()), &response)

	assert.Nil(t, err)
	assert.Equal(t, len(response), length)
	if length > 0 {
		assert.Equal(t, response[0].ID, 1)
		assert.Equal(t, response[0].Title, title)
		assert.Equal(t, response[0].IsDone, isdone)
	}
}

func doDeleteItem(r http.Handler, t *testing.T, id int) {
	request := doRequest(r, "DELETE", fmt.Sprintf("/api/%v", id), "")

	assert.Equal(t, http.StatusOK, request.Code)

	var response map[string]string

	err := json.Unmarshal([]byte(request.Body.String()), &response)
	value, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "OK", value)
}

func doPatchItem(r http.Handler, t *testing.T, id int, payload string) {
	request := doRequest(r, "PATCH", fmt.Sprintf("/api/%v", id), payload)

	assert.Equal(t, http.StatusOK, request.Code)

	var response map[string]string

	err := json.Unmarshal([]byte(request.Body.String()), &response)
	value, exists := response["message"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "OK", value)
}

func doRequest(r http.Handler, method string, path string, payload string) *httptest.ResponseRecorder {
	var req *http.Request

	if method == "POST" || method == "PATCH" || method == "PUT" {
		req, _ = http.NewRequest(method, path, bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, _ = http.NewRequest(method, path, nil)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}
