SHELL := /bin/zsh

dev:
	@echo "Starting..."
	cd frontend && npm run dev & \
	cd ../go && go run . & \
	wait