package dadJokeApi

import (
	"encoding/json"

	"fmt"
)

type DadJokeApiResponse struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func FromJson(rawData []byte) DadJokeApiResponse {
	res := DadJokeApiResponse{}

	if err := json.Unmarshal(rawData, &res); err != nil {
		fmt.Printf("Could not unmarshal raw data. %v", err)
	}

	return res
}

func ToJson(res DadJokeApiResponse) []byte {
	rawData, err := json.Marshal(res)

	if err != nil {
		fmt.Printf("Could not unmarshal raw data. %v", err)
		return []byte{}
	}

	return rawData
}
