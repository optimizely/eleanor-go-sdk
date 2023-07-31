// This file was auto-generated by Fern from our API Definition.

package api

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	core "github.com/fern-optimizely/optimizely-go/core"
	io "io"
	http "net/http"
)

type Client interface {
	Chat(ctx context.Context, request *ChatRequestDto) (*TokensResponseDto, error)
	ChatCompletion(ctx context.Context, request *ChatRequestDto) (*ChatCompletionsResponseDto, error)
	Proxy() ProxyClient
	Models() ModelsClient
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:      options.BaseURL,
		httpClient:   options.HTTPClient,
		header:       options.ToHeader(),
		proxyClient:  NewProxyClient(opts...),
		modelsClient: NewModelsClient(opts...),
	}
}

type client struct {
	baseURL      string
	httpClient   core.HTTPClient
	header       http.Header
	proxyClient  ProxyClient
	modelsClient ModelsClient
}

func (c *client) Chat(ctx context.Context, request *ChatRequestDto) (*TokensResponseDto, error) {
	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/tokens"

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

	var response *TokensResponseDto
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *client) ChatCompletion(ctx context.Context, request *ChatRequestDto) (*ChatCompletionsResponseDto, error) {
	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/complete"

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

	var response *ChatCompletionsResponseDto
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *client) Proxy() ProxyClient {
	return c.proxyClient
}

func (c *client) Models() ModelsClient {
	return c.modelsClient
}