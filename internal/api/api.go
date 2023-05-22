package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/samber/lo"
)

type EmoteSet struct {
	ID   string
	Name string
	// Emotes names by id
	Emotes map[string]string
}

type ApiEmoteSet interface {
	Create(name string) (emoteSetID string, err error)
	Read(emoteSetID string) (EmoteSet, error)
	UpdateName(emoteSetID, name string) (EmoteSet, error)
	Delete(emoteSetID string) error
}

type ApiEmote interface {
	AddToSet(emoteSetID, emoteID, emoteName string) error
	Read(emoteSetID, emoteID string) error
	UpdateName(emoteSetID, emoteID string, emoteName string) error
	RemoveFromSet(emoteSetID, emoteID string) error
}

type Api interface {
	EmoteSet() ApiEmoteSet
	Emote() ApiEmote
}

type api struct {
	apiEndpoint string
	token       string
	// TODO: http.Client
}

func NewClient(token string) (Api, error) {
	return api{
		apiEndpoint: "https://7tv.io/v3/gql",
		token:       token,
	}, nil
}

func (p api) apiCall(
	operationName string,
	variables any,
	query string,
	response any,
) error {
	payload := struct {
		OperationName string `json:"operationName"`
		Variables     any    `json:"variables"`
		Query         string `json:"query"`
	}{
		OperationName: operationName,
		Variables:     variables,
		Query:         query,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", p.apiEndpoint, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}

type apiEmoteSet api

func (p api) EmoteSet() ApiEmoteSet {
	return apiEmoteSet(p)
}

func (p apiEmoteSet) Create(name string) (string, error) {
	return "emote set id", errors.New("not implemented")
}

func (p apiEmoteSet) Read(emoteSetID string) (EmoteSet, error) {
	type Variables struct {
		ID string `json:"id"`
	}

	query := `query GetEmoteSet(
			$id: ObjectID!,
			$formats: [ImageFormat!]
		) {
			emoteSet(id: $id) {
				id
				name
				flags
				capacity
				origins {
					id
					weight
					__typename
				}
				emotes {
					id
					name
					actor {
						id
						display_name
						avatar_url
						__typename
						}
					origin_id
					data {
						id
						name
						flags
						state
						lifecycle
						host {
							url
							files(formats: $formats) {
								name
								format
								__typename
							}
							__typename
						}
						owner {
							id
							display_name
							style {
								color
								__typename
							}
							roles
							__typename
						}
						__typename
					}
					__typename
				}
				owner {
					id
					username
					display_name
					style {
						color
						__typename
					}
					avatar_url
					roles
					connections {
						id
						display_name
						emote_capacity
						platform
						__typename
					}
					__typename
				}
				__typename
			}
		}`

	type ResponseEmote struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Actor struct {
			ID          string `json:"id"`
			DisplayName string `json:"display_name"`
			AvatarURL   string `json:"avatar_url"`
			Typename    string `json:"__typename"`
		} `json:"actor"`
		OriginID interface{} `json:"origin_id"`
		Data     struct {
			ID        string   `json:"id"`
			Name      string   `json:"name"`
			Flags     int      `json:"flags"`
			States    []string `json:"state"`
			Lifecycle int      `json:"lifecycle"`
			Host      struct {
				URL   string `json:"url"`
				Files []struct {
					Name     string `json:"name"`
					Format   string `json:"format"`
					Typename string `json:"__typename"`
				} `json:"files"`
				Typename string `json:"__typename"`
			} `json:"host"`
			Owner struct {
				ID          string `json:"id"`
				DisplayName string `json:"display_name"`
				Style       struct {
					Color    int    `json:"color"`
					Typename string `json:"__typename"`
				} `json:"style"`
				Roles    []string `json:"roles"`
				Typename string   `json:"__typename"`
			} `json:"owner"`
			Typename string `json:"__typename"`
		} `json:"data"`
		Typename string `json:"__typename"`
	}
	type Response struct {
		Data struct {
			EmoteSet struct {
				ID       string          `json:"id"`
				Name     string          `json:"name"`
				Flags    int             `json:"flags"`
				Capacity int             `json:"capacity"`
				Origins  []interface{}   `json:"origins"`
				Emotes   []ResponseEmote `json:"emotes"`
				Owner    struct {
					ID          string `json:"id"`
					Username    string `json:"username"`
					DisplayName string `json:"display_name"`
					Style       struct {
						Color    int    `json:"color"`
						Typename string `json:"__typename"`
					} `json:"style"`
					AvatarURL   string   `json:"avatar_url"`
					Roles       []string `json:"roles"`
					Connections []struct {
						ID            string `json:"id"`
						DisplayName   string `json:"display_name"`
						EmoteCapacity int    `json:"emote_capacity"`
						Platform      string `json:"platform"`
						Typename      string `json:"__typename"`
					} `json:"connections"`
					Typename string `json:"__typename"`
				} `json:"owner"`
				Typename string `json:"__typename"`
			} `json:"emoteSet"`
		} `json:"data"`
	}

	var response Response

	if err := api(p).apiCall(
		"GetEmoteSet",
		Variables{
			ID: emoteSetID,
		},
		query,
		&response,
	); err != nil {
		return EmoteSet{}, err
	}

	return EmoteSet{
		ID:   response.Data.EmoteSet.ID,
		Name: response.Data.EmoteSet.Name,
		Emotes: lo.Associate(
			response.Data.EmoteSet.Emotes,
			func(emote ResponseEmote) (string, string) {
				return emote.Name, emote.ID
			},
		),
	}, nil
}

func (p apiEmoteSet) UpdateName(emoteSetID, name string) (EmoteSet, error) {
	type VariableData struct {
		Name     string  `json:"name"`
		Capacity int     `json:"capacity"`
		Origins  *string `json:"origins"`
	}
	type Variables struct {
		ID   string       `json:"id"`
		Data VariableData `json:"data"`
	}

	query := `mutation UpdateEmoteSet(
			$id: ObjectID!,
			$data: UpdateEmoteSetInput!
		) {
			emoteSet(id: $id) {
				update(data: $data) {
					id
					name
					__typename
				}
				__typename
			}
		}`

	type Response struct {
		Data struct {
			EmoteSet struct {
				Update struct {
					ID       string `json:"id"`
					Name     string `json:"name"`
					Typename string `json:"__typename"`
				} `json:"update"`
				Typename string `json:"__typename"`
			} `json:"emoteSet"`
		} `json:"data"`
	}

	var response Response

	if err := api(p).apiCall(
		"UpdateEmoteSet",
		Variables{
			ID: emoteSetID,
			Data: VariableData{
				Name:     name,
				Capacity: 300,
				Origins:  nil,
			},
		},
		query,
		&response,
	); err != nil {
		return EmoteSet{}, err
	}

	return EmoteSet{
		ID:     response.Data.EmoteSet.Update.ID,
		Name:   response.Data.EmoteSet.Update.Name,
		Emotes: nil,
	}, nil
}

func (p apiEmoteSet) Delete(emoteSetID string) error {
	return errors.New("not implemented")
}

type apiEmote api

func (p api) Emote() ApiEmote {
	return apiEmote(p)
}

func (p apiEmote) updateEmoteInSet(action, emoteSetID, emoteID, emoteName string) error {
	type Variables struct {
		Action     string `json:"action"` // "ADD",
		EmoteSetID string `json:"id"`
		EmoteID    string `json:"emote_id"` // "61801776e0801fb98788c028",
		Name       string `json:"name"`     // "MMMM"
	}

	query := `mutation ChangeEmoteInSet(
		$id: ObjectID!,
		$action: ListItemAction!,
		$emote_id: ObjectID!,
		$name: String
	) {
		emoteSet(id: $id) {
		    id
		    emotes(
				id: $emote_id,
				action: $action,
				name: $name
			) {
				id
				name
				__typename
			}
			__typename
		}
	}`

	// TODO: emote set returned
	var response map[string]any

	return api(p).apiCall(
		"ChangeEmoteInSet",
		Variables{
			Action:     action,
			EmoteSetID: emoteSetID,
			EmoteID:    emoteID,
			Name:       emoteName,
		},
		query,
		&response,
	)
}

func (p apiEmote) AddToSet(emoteSetID, emoteID, emoteName string) error {
	return p.updateEmoteInSet("ADD", emoteSetID, emoteID, emoteName)
}

func (p apiEmote) UpdateName(emoteSetID, emoteID, emoteName string) error {
	return p.updateEmoteInSet("UPDATE", emoteSetID, emoteID, emoteName)
}

func (p apiEmote) Read(emoteSetID, emoteID string) error {
	return errors.New("not implemented")
}

func (p apiEmote) RemoveFromSet(emoteSetID, emoteID string) error {
	// TODO: extract emote name for last arg?
	return p.updateEmoteInSet("REMOVE", emoteSetID, emoteID, "")
}
