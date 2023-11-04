build:
	export GOPROXY=https://goproxy.cn && go build -o bin/mobilequery .
start:	
	nohup /data/mobilequery/bin/mobilequery -b="0.0.0.0:8040" &
stop:
	killall mobilequery
test:
	curl "http://127.0.0.1:8040/query?mobile="
docker-image:
	DOCKER_BUILDKIT=1 docker build -t hetao29/mobilequery .
docker-image-nocache:
	DOCKER_BUILDKIT=1 docker build --no-cache -t hetao29/mobilequery .
docker-push:
	docker push hetao29/mobilequery:latest
