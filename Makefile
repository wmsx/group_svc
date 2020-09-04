.PHONY: docker
docker:
	docker build . -t group-svc:latest
