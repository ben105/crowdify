SHELL = /bin/sh

em:
	cd emulator && docker compose up --build

.PHONY: em