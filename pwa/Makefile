SERVICE=pwa
IMG_HUB?=docker.io/jmzwcn
TAG?=${shell date "+%Y%m%d-%H%M"}

clean:
	rm -rf dist
#	-cp resources/sign/* platforms/android

build:clean
#	ionic build --prod
	./node_modules/.bin/vue-cli-service build
#	docker build -t $(IMG_HUB)/$(SERVICE):$(TAG) .

image:
	docker build --platform=linux/arm64 -t $(IMG_HUB)/$(SERVICE):$(TAG) .

#push:build
push:image
	docker push $(IMG_HUB)/$(SERVICE):$(TAG)

resource:
	ionic cordova resources android
