package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rprtr258/log"
	"github.com/urfave/cli/v2"

	"github.com/rprtr258/seventv-tf-provider/internal/api"
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

func main() {
	if err := (&cli.App{
		Name:            "7tv",
		Usage:           "cli utility to manage emote sets",
		HideHelp:        true,
		HideHelpCommand: true,
		HideVersion:     true,
		Commands: []*cli.Command{
			{
				Name:  "push",
				Usage: "push emotes state from config to 7tv",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "token",
						Usage: "auth token, get from requests to gql api on 7tv.app",
					},
					&cli.StringFlag{
						Name:      "config",
						Usage:     "config file with emotesets descriptions",
						TakesFile: true,
					},
				},
				Action: func(ctx *cli.Context) error {
					token := ctx.String("token")
					configPath := ctx.String("config")

					configBytes, err := os.ReadFile(configPath)
					if err != nil {
						return fmt.Errorf("read config %s: %w", configPath, err)
					}

					var config map[string]struct {
						Name   string            `json:"name"`
						Emotes map[string]string `json:"emotes"`
					}
					if err := json.Unmarshal(configBytes, &config); err != nil {
						return fmt.Errorf("parse config: %w", err)
					}

					client, err := api.NewClient(token)
					if err != nil {
						return fmt.Errorf("new client: %w", err)
					}

					for emotesetID, emoteset := range config {
						actualEmoteset, err := client.EmoteSet().Read(emotesetID)
						if err != nil {
							return fmt.Errorf("get setID=%s: %w", emotesetID, err)
						}

						toRemove, toAdd := mapDiff(actualEmoteset.Emotes, emoteset.Emotes)

						for name, emoteID := range toRemove {
							log.Infof("removing emote", log.F{"name": name, "emoteID": emoteID})
							if err := client.Emote().RemoveFromSet(emotesetID, emoteID); err != nil {
								return fmt.Errorf("remove emoteID=%q name=%q: %w", emoteID, name, err)
							}
						}

						for name, emoteID := range toAdd {
							log.Infof("adding emote", log.F{"name": name, "emoteID": emoteID})
							if err := client.Emote().AddToSet(emotesetID, emoteID, name); err != nil {
								return fmt.Errorf(
									"add emoteID=%s name=%s to setID=%s: %w",
									emoteID, name, emotesetID, err,
								)
							}
						}
					}

					return nil

				},
			},
		},
	}).Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
