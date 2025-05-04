EVENT_COLLECTOR=eventCollector
PROFILE_MONITOR=profileMonitor

up:
	@echo Starting Docker images...
	docker compose up -d
	@echo Docker Images started!

up_build: build_event_collector build_profile_monitor
	@echo Stopping docker images \(if running...\)
	docker compose down
	@echo Building \(when requeired\) and starting docker images...
	docker compose up --build -d
	@echo Docker images built and started!

down:
	@echo Stopping docker compose...
	docker compose down
	@echo Done!

build_event_collector:
	@echo Building eventCollector binary
	chdir ../EventCollector && set GOOS=linux && set GOARCH=amd64 && set CGO_ENABLED=0 && go build -o ${EVENT_COLLECTOR} ./cmd/api
	@echo Done!

build_profile_monitor:
	@echo Building profileMonitor binary
	chdir ../ProfileMonitor && set GOOS=linux && set GOARCH=amd64 && CGO_ENABLED=0 && go build -o ${PROFILE_MONITOR} ./cmd/api
	@echo Done!