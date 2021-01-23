package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func httpGet(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("HTTP error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("expected status OK but got %q instead", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read error: %s\n", err)
	}

	return string(data), nil
}

func httpPost(url string) (string, error) {
	payload := "Hello world!"
	resp, err := http.Post(url, "text/plain", strings.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("HTTP error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("expected status OK but got %q instead", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read error: %s\n", err)
	}

	return string(data), nil
}

func getUserAgent(jsonData string) (string, error) {
	type response struct {
		Headers struct {
			UserAgent string `json:"User-Agent"`
		} `json:"headers"`
	}

	var resp response
	err := json.Unmarshal([]byte(jsonData), &resp)
	if err != nil {
		return "", fmt.Errorf("json parsing error: %s\n", err)
	}

	return resp.Headers.UserAgent, nil
}

func getWithTimeout(url string, timeout time.Duration) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("envalid URL: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return "", fmt.Errorf("HTTP error: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("expected status OK but got %q instead", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read error: %s\n", err)
	}

	return string(data), nil
}

func main() {
	//////////////
	// HTTP GET //
	//////////////

	payload, err := httpGet("https://httpbin.org/get")
	if err != nil {
		fmt.Printf("GET error: %s\n", err)
	}

	fmt.Println("GET response:")
	fmt.Println(payload)

	///////////////
	// HTTP POST //
	///////////////

	payload, err = httpPost("https://httpbin.org/post")
	if err != nil {
		fmt.Printf("POST error: %s\n", err)
	}

	fmt.Println("POST response:")
	fmt.Println(payload)

	////////////////////
	// Get User Agent //
	////////////////////

	payload, err = httpGet("https://httpbin.org/headers")
	if err != nil {
		fmt.Printf("GET error: %s\n", err)
	}

	userAgent, err := getUserAgent(payload)
	if err != nil {
		fmt.Printf("Failed to extract user agent: %s\n", err)
	}

	fmt.Printf("Request user agent: %s\n\n", userAgent)

	//////////////////////
	// Get with timeout //
	//////////////////////

	httpbinTimeout := 5
	httpTimeout := 10 * time.Second
	payload, err = getWithTimeout(
		fmt.Sprintf("https://httpbin.org/delay/%d", httpbinTimeout),
		httpTimeout,
	)
	if err != nil {
		fmt.Printf("GET with timeout error: %s\n", err)
	}

	fmt.Println("GET with timeout response:")
	fmt.Println(payload)
}
