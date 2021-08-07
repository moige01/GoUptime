// Copyright (c) 2021 Moisés González
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type DiscordHandler struct {
	token   string
	channel string
}

func (d *DiscordHandler) init() {
	token, ok := os.LookupEnv("DISCORD_BOT_TOKEN")

	if !ok {
		ErrorLogger.Println("DISCORD_TOKEN is not set!")

		return
	}

	channel, ok := os.LookupEnv("DISCORD_CHANNEL_ID")

	if !ok {
		ErrorLogger.Println("DISCORD_TOKEN is not set!")
	}

	d.token = token
	d.channel = channel
}

func (d *DiscordHandler) Dispatch(page string, status int, message string, wg *sync.WaitGroup) {
	defer wg.Done()

	d.init()

	dg, err := discordgo.New("Bot " + d.token)

	defer dg.Close()

	if err != nil {
		ErrorLogger.Println("error creating Discord session,", err)

		return
	}

	if status >= 200 && status < 300 {
		return
	}

	_, err = dg.ChannelMessageSend(d.channel, fmt.Sprintf("[%s] %d - %s", page, status, message))

	if err != nil {
		ErrorLogger.Println("error sending Discord message,", err)
	}
}
