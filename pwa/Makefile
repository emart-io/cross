SERVICE=pwa
IMG_HUB?=docker.io/jmzwcn
TAG?=${shell date "+%Y%m%d-%H%M"}

clean:
	rm -rf dist
#	-cp resources/sign/* platforms/android

build:clean
#	ionic build --prod
	./node_modules/.bin/vue-cli-service build --no-module
#	docker build -t $(IMG_HUB)/$(SERVICE):$(TAG) .

push:build
	docker build --platform=linux/arm64 -t $(IMG_HUB)/$(SERVICE):$(TAG) .
	docker push $(IMG_HUB)/$(SERVICE):$(TAG)

resource:
	ionic cordova resources android
