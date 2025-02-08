.PHONY: test

test:
	docker compose run --rm backend go test ./... -v

