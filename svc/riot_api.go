package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/BenjixD/bengi-lol-bot/utils/reflections"
	"github.com/go-resty/resty/v2"
)

const RiotGamesUSApiUrl = "https://americas.api.riotgames.com"

type GetUserFromRiotIDRequest struct {
	GameName string
	TagLine  string
}

type GetUserFromRiotIDResponse struct {
	Puuid string
}

type GetMatchesByPuuidRequest struct {
	Puuid  string
	_Query struct {
		StartTime *string `optional:"true"`
		EndTime   *string `optional:"true"`
		Queue     *string `optional:"true"`
		Start     *string `optional:"true"`
		Count     *string `optional:"true"`
	}
}

type GetMatchesByPuuidResponse []string

type RiotApiClientInterface interface {
	GetUserFromRiotID(ctx context.Context, req *GetUserFromRiotIDRequest) (*GetUserFromRiotIDResponse, error)
	GetMatchesByPuuid(ctx context.Context, req *GetMatchesByPuuidRequest) (*GetMatchesByPuuidResponse, error)
}

type RiotApiClient struct {
	client *resty.Client
}

func NewRiotApiClient(apiKey string) (*RiotApiClient, error) {
	return &RiotApiClient{
		client: NewApiClient(
			RiotGamesUSApiUrl,
			map[string]string{
				"X-Riot-Token": apiKey,
				"Content-Type": "application/json",
			}),
	}, nil
}

func (c *RiotApiClient) GetUserFromRiotID(ctx context.Context, req *GetUserFromRiotIDRequest) (*GetUserFromRiotIDResponse, error) {
	path := fmt.Sprintf("riot/account/v1/accounts/by-riot-id/%s/%s", req.GameName, req.TagLine)
	sanitizedPath, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	res, err := c.client.R().Get(sanitizedPath.Path)
	if err != nil {
		return nil, err
	}

	if res.IsSuccess() {
		var apiRes GetUserFromRiotIDResponse
		err := json.Unmarshal(res.Body(), &apiRes)
		if err != nil {
			return nil, err
		}
		return &apiRes, nil
	}
	return nil, fmt.Errorf("unexpected code %d", res.StatusCode())
}

func (c *RiotApiClient) GetMatchesByPuuid(ctx context.Context, req *GetMatchesByPuuidRequest) (*GetMatchesByPuuidResponse, error) {
	path := fmt.Sprintf("lol/match/v5/matches/by-puuid/%s/ids", req.Puuid)
	sanitizedPath, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	var queryParams map[string]string
	for k, v := range reflections.StructToMap(req._Query) {
		// Note: Since Query Params are either str* or str, we need
		// to get the string value referenced by the ptr, not the actual
		// pointer value
		if reflections.IsStringPtr(v) {
			vString := v.(*string)
			queryParams[k] = fmt.Sprintf("%v", *vString)
		} else {
			queryParams[k] = fmt.Sprintf("%v", v)
		}
	}

	res, err := c.client.R().Get(sanitizedPath.Path)
	c.client.R().SetQueryParams(queryParams)
	if err != nil {
		return nil, err
	}

	if res.IsSuccess() {
		var apiRes GetMatchesByPuuidResponse
		err := json.Unmarshal(res.Body(), &apiRes)
		if err != nil {
			return nil, err
		}
		return &apiRes, nil
	}
	return nil, fmt.Errorf("unexpected code %d", res.StatusCode())
}
