run-mongo:
	docker run -d -p 27017:27017 --name mongodb mongo:4.4

stop-mongo:
	docker stop mongodb
	docker rm mongodb
