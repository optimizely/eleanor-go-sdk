// This file was auto-generated by Fern from our API Definition.

package api

type ChatRequestDto struct {
	Model    string            `json:"model"`
	Messages []*ChatMessageDto `json:"messages,omitempty"`
}
