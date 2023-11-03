docker-build:
	docker build --tag httpserv .

docker-run:
	docker run -it --rm \
	--name http-server \
	-p 8080:8080 \
	httpserv


docker-tag:
	docker tag httpserv europe-west10-docker.pkg.dev/celestial-geode-404007/test-repo/httpserv

docker-push:
	docker push europe-west10-docker.pkg.dev/celestial-geode-404007/test-repo/httpserv:latest

heroku-stack:
	heroku stack:set container --app joshsteveth-test