// Copyright (C) 2021  Antonio Lassandro

// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.

// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
// more details.

// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func run(authtoken string) int {
	baseURL := fmt.Sprint("https://api.telegram.org/bot", authtoken, "/")

	updateOffset := 0

	for true {
		resp, err := http.PostForm(
			fmt.Sprint(baseURL, "getUpdates"), url.Values{
				"offset": []string{strconv.Itoa(updateOffset)},
				"timeout": []string{"1"},
				"allowed_updates": []string{"[message]"},
			},
		)

		if err != nil {
			log.Println("Could not get Telegram updates")
			continue
		}

		if resp.StatusCode != http.StatusOK {
			log.Println(fmt.Sprint("Received bad response: ", resp.Status))
			continue
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Println("Could not read Telegram updates")
			continue
		}

		jsonResp := make(map[string]interface{})
		if err := json.Unmarshal(body, &jsonResp); err != nil {
			log.Println("Could not read Telegram updates")
			continue
		}

		if apiOk, ok := jsonResp["ok"].(bool); !ok {
			log.Println("Missing 'ok' in response")
			continue
		} else if !apiOk {
			if apiDescription, ok := jsonResp["description"].(string); ok {
				log.Println(fmt.Sprint("API error:", apiDescription))
				continue
			} else {
				log.Println("Unknown API error")
				continue
			}
		}

		results, ok := jsonResp["result"].([]interface{})
		if !ok {
			log.Println("Missing 'result' in response")
		}

		for _, result := range results {
			update, _ := result.(map[string]interface{})

			updateId, ok := update["update_id"].(float64)
			if !ok {
				log.Println("Received updates without IDs")
				continue
			}

			updateOffset = int(updateId) + 1

			message, ok := update["message"].(map[string]interface{})
			if !ok {
				continue
			}

			text, ok := message["text"].(string)
			if !ok {
				continue
			}

			if !strings.HasPrefix(text, "/hand") {
				continue
			}

			chat, ok := message["chat"].(map[string]interface{})
			if !ok {
				log.Println("Missing 'chat' from message")
				continue
			}

			chatId, ok := chat["id"].(float64)
			if !ok {
				log.Println("Missing 'chat.id' from message")
				continue
			}

			resp, err := http.PostForm(
				fmt.Sprint(baseURL, "sendMessage"), url.Values{
					"chat_id": []string{strconv.Itoa(int(chatId))},
					"text": []string{MakeHand()},
				},
			)

			if err != nil {
				log.Println("Could not send Telegram message")
				continue
			}

			if resp.StatusCode != http.StatusOK {
				log.Println("Could not send Telegram message")
				continue
			}
		}
	}

	return 0
}

func main() {
	rand.Seed(time.Now().UnixNano())

	authtoken := flag.String("token", "", "The bot's auth token")
	flag.Parse()

	if len(*authtoken) == 0 {
		log.Fatal("Auth token not provided")
	}

	log.SetOutput(os.Stdout)

	os.Exit(run(strings.TrimSpace(*authtoken)))
}
