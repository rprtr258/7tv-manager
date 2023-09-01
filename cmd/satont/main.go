package main

import (
	"fmt"
	"os"

	"github.com/rprtr258/log"
	"github.com/urfave/cli/v2"

	"github.com/rprtr258/7tv-manager/internal/api"
)

func mapDiff(old, new map[string]string) (toRemove, toAdd map[string]string) {
	toRemove = map[string]string{}
	for name, oldEmoteID := range old {
		newEmoteID, ok := new[name]
		if !ok || oldEmoteID != newEmoteID {
			toRemove[name] = oldEmoteID
		}
	}

	toAdd = map[string]string{}
	for name, newEmoteID := range new {
		oldEmoteID, ok := old[name]
		if !ok || oldEmoteID != newEmoteID {
			toAdd[name] = newEmoteID
		}
	}

	return
}

type EmotesetConfig struct {
	Name   string            `json:"name"`
	Emotes map[string]string `json:"emotes"`
}

type Config map[string]EmotesetConfig

var _app = &cli.App{
	Name:            "7tv-satont",
	Usage:           "cli utility to transfer emote sets blazingly fast",
	HideHelp:        true,
	HideHelpCommand: true,
	HideVersion:     true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "token",
			Usage: "auth token, get from requests to gql api on 7tv.app",
		},
		&cli.StringFlag{
			Name:  "from",
			Usage: "emote set id to pull (for now it cannot pull all emotesets)",
		},
		&cli.StringFlag{
			Name:  "to",
			Usage: "emote set id to push to",
		},
	},
	Action: func(ctx *cli.Context) error {
		token := ctx.String("token")
		emoteSetIDFrom := ctx.String("from")
		emoteSetIDTo := ctx.String("to")

		client, err := api.NewClient(token)
		if err != nil {
			return fmt.Errorf("new client: %w", err)
		}

		emotesetFrom, err := client.EmoteSet().Read(emoteSetIDFrom)
		if err != nil {
			return fmt.Errorf("get setID=%s: %w", emoteSetIDFrom, err)
		}

		for name, emoteID := range emotesetFrom.Emotes {
			log.Infof("adding emote", log.F{"name": name, "emoteID": emoteID})
			if err := client.Emote().AddToSet(emoteSetIDTo, emoteID, name); err != nil {
				return fmt.Errorf(
					"add emoteID=%s name=%s to setID=%s: %w",
					emoteID, name, emoteSetIDTo, err,
				)
			}
		}

		return nil
	},
}

func main() {
	if err := _app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
