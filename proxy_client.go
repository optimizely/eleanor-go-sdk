// This file was auto-generated by Fern from our API Definition.

package api

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	core "github.com/fern-optimizely/optimizely-go/core"
	io "io"
	http "net/http"
)

type ProxyClient interface {
	AzureOpenai(ctx context.Context, path string, request map[string]any) (any, error)
	Openai(ctx context.Context, path string, request map[string]any) (any, error)
}

func NewProxyClient(opts ...core.ClientOption) ProxyClient {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &proxyClient{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type proxyClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// Proxy requests to Azure OpenAI REST API path /openai/{path},
// see [Azure OpenAI API reference](https://learn.microsoft.com/en-us/azure/cognitive-services/openai/reference).
//
// **Note:** The request is forwarded to a random backend, some deployments are not available in all backends.
func (p *proxyClient) AzureOpenai(ctx context.Context, path string, request map[string]any) (any, error) {
	baseURL := ""
	if p.baseURL != "" {
		baseURL = p.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/raw/openai/%v", path)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 422:
			value := new(UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response any
	if err := core.DoRequest(
		ctx,
		p.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		p.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

// Proxy requests to Azure OpenAI REST API path /openai/{path},
// see [Azure OpenAI API reference](https://learn.microsoft.com/en-us/azure/cognitive-services/openai/reference).
//
// **Note:** The request is forwarded to a random backend, some deployments are not available in all backends.
func (p *proxyClient) Openai(ctx context.Context, path string, request map[string]any) (any, error) {
	baseURL := ""
	if p.baseURL != "" {
		baseURL = p.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/raw/openai/%v", path)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 422:
			value := new(UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response any
	if err := core.DoRequest(
		ctx,
		p.httpClient,
		endpointURL,
		http.MethodPut,
		request,
		&response,
		false,
		p.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}
