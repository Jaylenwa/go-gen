add:
	gofmt -s -w .
	git add .

push:
	git pull origin $(shell git rev-parse --abbrev-ref HEAD)
	git push origin $(shell git rev-parse --abbrev-ref HEAD)

pull:
	git pull origin $(shell git rev-parse --abbrev-ref HEAD)