# go-chat-server
Basic Chat Server

# Get Repo
`go get github.com/rybitskyi/go-chat-server`

# Go version
Go 1.9+

# APIs

## Check Status

```
curl -H "Content-Type: application/json" http://localhost:8081/status
```

## Get Messages

```
curl -H "Content-Type: application/json" http://localhost:8081/messages
```

## Get Users

```
example:
curl -H "Content-Type: application/json" http://localhost:8081/users
```

# Future Improvements
## [POST /message] Return created object
## [POST /message] Return exception error