/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"

	"internal/dadJokeApi"
)

const baseApiUrl string = "https://icanhazdadjoke.com/"

var customHeaders map[string]string = map[string]string{"User-Agent": "Dad Joke CLI for learning Go (https://github.com/willowell/dadjoke-cli)"}

var myClient = &http.Client{Timeout: 10 * time.Second}

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke",
	Long:  "Fetch a random dad joke from the icanhazdadjoke API",
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

func getRandomJoke() {
	responseBytes := getJson(myClient, baseApiUrl, customHeaders)

	decodedRes := dadJokeApi.FromJson(responseBytes)

	fmt.Println(decodedRes.Joke)
}

func getJson(client *http.Client, baseApiUrl string, customHeaders map[string]string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseApiUrl,
		nil,
	)

	if err != nil {
		log.Printf("Could not construct a request. %v\n", err)
		return []byte{}
	}

	request.Header.Add("Accept", "application/json")

	if len(customHeaders) > 0 {
		for header, value := range customHeaders {
			request.Header.Add(header, value)
		}
	}

	response, err := client.Do(request)

	if err != nil {
		log.Printf("Could not make a request. %v\n", err)
		return []byte{}
	}

	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("Could not read response body. %v\n", err)
		return []byte{}
	}

	return responseBytes
}
