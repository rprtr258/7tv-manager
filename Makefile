TOKEN := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiNjNiNTgwMmFlYmZiYzJkYzRkZjI4ODQ3IiwidiI6MCwiaXNzIjoic2V2ZW50di1hcGkiLCJleHAiOjE2OTI1NDMwOTYsIm5iZiI6MTY4NDc2NzA5NiwiaWF0IjoxNjg0NzY3MDk2fQ.DCW35XYMV9LvsV_0YWvofQjJsaQ28Qijw7qeePLVjT8"
EMOTESET_ID := "63b58083c032521d3d256191"
7TV := go run cmd/7tv/main.go

push:
	@${7TV} push --token ${TOKEN} --config examples/push.jsonnet

pull:
	@${7TV} pull --token ${TOKEN} --config examples/pull.jsonnet --id ${EMOTESET_ID}
