package utils

import (
	"fmt"
	"strings"
)

// BidRequest represents the structure of a bid request payload.
type BidRequest struct {
	Id     string `json:"id" binding:"required"`
	Width  *int   `json:"width,omitempty"`
	Height *int   `json:"height,omitempty"`
	Banner struct {
		Type int `json:"type" binding:"required"`
	} `json:"banner"`
}

// ValidateBidRequest validates the bid request payload.
func ValidateBidRequest(request BidRequest) error {
	if request.Banner.Type != 1 && request.Banner.Type != 2 {
		return fmt.Errorf("banner type must be 1 or 2")
	}
	if (request.Width != nil && request.Height == nil) || (request.Width == nil && request.Height != nil) {
		return fmt.Errorf("both width and height must be provided together")
	}
	return nil
}

// ReplacePlaceholders replaces placeholders in the template with actual data.
func ReplacePlaceholders(template string, data map[string]string) string {
	for key, value := range data {
		placeholder := fmt.Sprintf("{%s}", key)
		template = strings.ReplaceAll(template, placeholder, value)
	}
	return template
}
