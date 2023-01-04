package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Emote struct {
	ID   string
	Name string
}

type EmoteSet struct {
	ID     string
	Name   string
	Emotes []Emote
}

type Api interface {
	CreateEmoteSet(name string) (emoteSetID string, err error)
	GetEmoteSet(emoteSetID string) (EmoteSet, error)
	UpdateEmoteSet(emoteSetID, name string) (EmoteSet, error)
	DeleteEmoteSet(emoteSetID string) error

	CreateEmoteBinding(emoteSetID, emoteID string, emoteName *string) error
	GetEmoteBinding(emoteSetID, emoteID string) error
	// UpdateEmoteBinding(emoteSetID, emoteID string, emoteName *string) error
	DeleteEmoteBinding(emoteSetID, emoteID string) error

	// GetEmote(emoteID string) (Emote, error)
}

type api struct {
	apiEndpoint string
	token       string
}

func NewClient(username, password *string) (Api, error) {
	return &api{
		apiEndpoint: "https://7tv.io/v3/gql",
		token:       "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiNjNiNTgwMmFlYmZiYzJkYzRkZjI4ODQ3IiwidiI6MCwiaXNzIjoiN1RWLUFQSS1SRVNUIiwiZXhwIjoxNjgwNzIzMjEwfQ.nnQROcsz1Wjlmzfloin2qaihdeZnT9kpfyIp9tx9izw",
	}, nil
}

func (p *api) CreateEmoteSet(name string) (string, error) {
	return "63b58083c032521d3d256191", nil
}

func (p *api) GetEmoteSet(emoteSetID string) (EmoteSet, error) {
	return EmoteSet{}, errors.New("not implemented")
}

func (p *api) UpdateEmoteSet(emoteSetID, name string) (EmoteSet, error) {
	return EmoteSet{}, errors.New("not implemented")
}

func (p *api) DeleteEmoteSet(emoteSetID string) error {
	return errors.New("not implemented")
}

func (p *api) changeEmoteInSet(action, emoteSetID, emoteID, emoteName string) error {
	bts, err := json.Marshal(
		map[string]any{
			"operationName": "ChangeEmoteInSet",
			"variables": map[string]string{
				"action":   action, // "ADD",
				"id":       emoteSetID,
				"emote_id": emoteID,   // "61801776e0801fb98788c028",
				"name":     emoteName, // "MMMM"
			},
			"query": `mutation ChangeEmoteInSet(
				$id: ObjectID!,
				$action: ListItemAction!,
				$emote_id: ObjectID!,
				$name: String
			) {
				emoteSet(id: $id) {
				    id
				    emotes(id: $emote_id, action: $action, name: $name) {
						id
						name
						__typename
					}
					__typename
				}
			}`,
		},
	)
	if err != nil {
		return err
	}

	body := bytes.NewReader(bts)

	_, err = http.NewRequest("POST", p.apiEndpoint, body)
	return err
}

func (p *api) CreateEmoteBinding(emoteSetID, emoteID string, emoteName *string) error {
	if emoteName == nil {
		// TODO: extract emote name
		return errors.New("not implemented")
	}

	return p.changeEmoteInSet("ADD", emoteSetID, emoteID, *emoteName)
}

func (p *api) GetEmoteBinding(emoteSetID, emoteID string) error {
	return errors.New("not implemented")
}

func (p *api) DeleteEmoteBinding(emoteSetID, emoteID string) error {
	// TODO: extract emote name for last arg?
	return p.changeEmoteInSet("DELETE", emoteSetID, emoteID, "")
}
