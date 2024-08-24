package models

import (
	"fmt"
	"go-api/config"
	"go-api/utils"
)

// IsIdPresent checks if a given ID is present in Redis.
func IsIdPresent(id string) bool {
	exists, err := config.RedisClient.Exists(config.Ctx, id).Result()
	if err != nil || exists == 0 {
		return false
	}
	return true
}

// GetResponseData fetches data based on request and returns the formatted response and content type.
func GetResponseData(request utils.BidRequest) (string, string, error) {
	data, err := config.RedisClient.HGetAll(config.Ctx, request.Id).Result()
	if err != nil || len(data) == 0 {
		return "", "", fmt.Errorf("failed to fetch data for ID %s", request.Id)
	}

	var template string
	var contentType string

	if request.Banner.Type == 1 {
		template, err = config.RedisClient.Get(config.Ctx, request.Id+"_js").Result()
		contentType = "text/javascript"
	} else if request.Banner.Type == 2 {
		template, err = config.RedisClient.Get(config.Ctx, request.Id+"_xml").Result()
		contentType = "text/xml"
	}

	if err != nil {
		return "", "", fmt.Errorf("failed to fetch template for ID %s", request.Id)
	}

	response := utils.ReplacePlaceholders(template, data)
	return response, contentType, nil
}
