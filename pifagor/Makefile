.PHONY:run
run: 
	sudo docker run -p 8080:8081 -it pifagor-app

.PHONY:build
build:
	sudo docker build -t pifagor-app .

.DEFAULT_GOAL := run