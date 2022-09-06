####################################################
.PHONY: botv1
botv1:
	go run ./cmd/v1/bot

####################################################
.PHONY: servicev1
servicev1:
	go run ./cmd/v1/service

####################################################
.PHONY: botv2
botv2:
	go run ./cmd/v2/bot

####################################################
.PHONY: servicev2
servicev2:
	go run ./cmd/v2/service

####################################################
MIGRATIONS_DIR=./migrations
.PHONY: migration
migration:
	goose -dir=${MIGRATIONS_DIR} create $(NAME) sql

####################################################
.PHONY: generate
generate:
	go generate ./...

####################################################
.PHONY: test
test:
	go test ./...

####################################################
.PHONY: coverage
coverage:
	go test -v $$(go list ./... | grep -v -E '/pkg/(api)') -covermode=count -coverprofile=/tmp/go_cover.out
	go tool cover -html=/tmp/go_cover.out
