
image:
	docker build --tag hainingzhen/alpine-mem-test:latest .
.PHONY: image

run:
	docker run -dp 6060:6060 --memory="128m" hainingzhen/alpine-mem-test:latest
.PHONY: run

push:
	docker push hainingzhen/alpine-mem-test:latest
.PHONY: docker-pushake

cluster:
	kind create cluster --config config.yaml
	kubectl config use-context kind-kind
.PHONY: cluster

