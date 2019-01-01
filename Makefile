build:
	go mod vendor
	env GOOS=linux go build -ldflags="-s -w" -o bin/slack-notify lambdas/notify/main.go