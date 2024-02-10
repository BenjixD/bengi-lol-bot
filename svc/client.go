package svc

import (
	"github.com/go-resty/resty/v2"
)

const ApiRetryCount = 3
const RetryWaitTimeSec = 1

func NewApiClient(baseUrl string, headers map[string]string) *resty.Client {
	client := resty.New()
	// Set retry configuration
	client.SetRetryCount(ApiRetryCount)
	client.SetRetryWaitTime(client.RetryWaitTime)
	// Base Url
	client.SetBaseURL(baseUrl)
	// default headers
	client.SetHeaders(headers)
	return client
}
