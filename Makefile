SHELL = /bin/sh

em:
	cd emulator && docker compose up --build

int:
	cd services/crowdify/integration && go test -count=1 ./...

.PHONY: em, int