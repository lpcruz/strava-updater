include .env

REPO_NAME=https://github.com/lpcruz/strava-updater
SRC=strava-updater
AUTH_URL=https://www.strava.com/oauth/authorize?client_id=${CLIENT_ID}&response_type=code&redirect_uri=http://localhost&approval_prompt=force&scope=activity:read_all,activity:write

ifndef GOPATH
export GOPATH=$(shell go env "GOPATH")
endif

ifndef CLIENT_CODE
$(info 🚨 CLIENT_CODE not set. Go to ${AUTH_URL} )
exit 1
endif

define parse_access_token
	grep -o '"access_token": *"[^"]*"' oauth_response.json | grep -o '"[^"]*"'
endef

fmt:
	gofmt -s -w ${SRC}

strava-update:
	STRAVA_ACCESS_TOKEN=${STRAVA_ACCESS_TOKEN} APP_NAME=${APP_NAME} go run strava-updater/main.go

strava-get-access-token:
	curl -X POST https://www.strava.com/oauth/token \
        -F client_id=${CLIENT_ID} \
        -F client_secret=${CLIENT_SECRET} \
        -F code=${CLIENT_CODE} \
        -F grant_type=authorization_code > oauth_response.json
	$(parse_access_token)


open-strava-auth:
	@if [ -z "$(CLIENT_ID)" ]; then \
		echo "ERROR: client_id is not set. Please provide it like: make open-strava-auth client_id=YOUR_CLIENT_ID"; \
		exit 1; \
	fi
	@URL="https://www.strava.com/oauth/authorize?client_id=$(CLIENT_ID)&response_type=code&redirect_uri=http://localhost&approval_prompt=force&scope=read,activity:write,activity:read"; \
	echo "Opening: $$URL"; \
	if command -v xdg-open >/dev/null; then \
		xdg-open "$$URL"; \
	elif command -v open >/dev/null; then \
		open "$$URL"; \
	else \
		echo "No known browser open command found (xdg-open or open). Please open manually: $$URL"; \
	fi
