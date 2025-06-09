.PHONY: run\
api\
model\
init_ddl\
build\
build_image\
save_image

IMAGE_NAME = deepkg_be
IMAGE_VERSION = 1.0.0
IMAGE_TAG = mickeyzhoudocker/${IMAGE_NAME}:${IMAGE_VERSION}

run:
	go run admin.go -f etc/admin.yaml

api:
	goctl api go -api api/main.api -dir ./

model:
	goctl model mysql ddl -src ./model/schema/ddl.sql -dir ./model/gozero_model

init_ddl:
	go run ./cmd/model/init.go -f ./etc/admin.yaml

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server admin.go

build_image:
	docker build -f Dockerfile --network=host --platform linux/amd64 -t ${IMAGE_TAG} .

save_image:
	docker save -o ${IMAGE_NAME}.tar ${IMAGE_TAG}