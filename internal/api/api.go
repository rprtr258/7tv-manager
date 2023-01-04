package api

// curl 'https://7tv.io/v3/gql' -X POST -H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0' -H 'Accept: */*' -H 'Accept-Language: en-US,en;q=0.5' -H 'Accept-Encoding: gzip, deflate, br' -H 'Referer: https://7tv.app/' -H 'content-type: application/json' -H 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiNjNiNTgwMmFlYmZiYzJkYzRkZjI4ODQ3IiwidiI6MCwiaXNzIjoiN1RWLUFQSS1SRVNUIiwiZXhwIjoxNjgwNzIzMjEwfQ.nnQROcsz1Wjlmzfloin2qaihdeZnT9kpfyIp9tx9izw' -H 'Origin: https://7tv.app' -H 'DNT: 1' -H 'Connection: keep-alive' -H 'Sec-Fetch-Dest: empty' -H 'Sec-Fetch-Mode: cors' -H 'Sec-Fetch-Site: cross-site' -H 'Pragma: no-cache' -H 'Cache-Control: no-cache' -H 'TE: trailers' --data-raw '{"operationName":"ChangeEmoteInSet","variables":{"action":"ADD","id":"63b58083c032521d3d256191","emote_id":"61801776e0801fb98788c028","name":"MMMM"},"query":"mutation ChangeEmoteInSet($id: ObjectID!, $action: ListItemAction!, $emote_id: ObjectID!, $name: String) {\n  emoteSet(id: $id) {\n    id\n    emotes(id: $emote_id, action: $action, name: $name) {\n      id\n      name\n      __typename\n    }\n    __typename\n  }\n}"}'

type Api interface {
	GetEmoteSet(emoteSetID string) (EmoteSet, error)
	GetEmote(emoteID string) (Emote, error)
}

func NewClient(username, password *string) (Api, error) {
	return nil, nil
}

type Emote struct {
	ID   string
	Name string
}

type EmoteSet struct {
	ID     string
	Emotes []Emote
}
