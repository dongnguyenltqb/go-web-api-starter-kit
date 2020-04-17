bkname = ./backup/ganja-$(shell date +"%Y-%m-%d-%H-%M")
pid = $(shell cat .pid +"%T")
logfile = log/$(pid).log

default: build

build:
	@clear
	@mkdir -p bin
	@echo "STEP : BUILD"
	@go build -o bin/ganja
	@echo "Build successfully."

install:
	@clear
	@echo "STEP : INSTALL"
	@go install
	@echo "Install successfully."

back-up:
	@echo "STEP : BACK UP"
	@mkdir -p ./backup
	@cp ./bin/ganja $(bkname)
	@echo "Done"

server:
	@clear
	@echo "STEP : RUN SERVER"
	@./bin/ganja server

stop:
	@clear
	@echo "Kill PID = $(pid)"
	@kill $(pid)
	@echo "=> done.."
