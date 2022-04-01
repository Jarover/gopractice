APP?=gopractice
RELEASE?=$(shell python version.py get)
COMMIT?=$(shell git rev-parse --short HEAD)
#BUILD_TIME?=$(shell powershell get-date -format "{yyyy-mm-dd_HH:mm:ss}")
BUILD_TIME?=$(shell date -u '+%Y-%m-%dT%H:%M:%S')
PROJECT?=github.com/Jarover/gopractice

clean:
	rm -f ${APP}
	rm -f ${APP}.exe


buildwin: clean
	python version.py inc-patch
	GOOS=windows go build \
				-o ${APP}.exe \
                -ldflags "-s -w \
				-X  ${PROJECT}/internal/app/config.Release=${RELEASE} \
                -X ${PROJECT}/internal/app/config.Commit=${COMMIT} \
				-X ${PROJECT}/internal/app/config.BuildTime=${BUILD_TIME}" \
                cmd/${APP}/main.go


buildlinux:	clean
	python version.py inc-patch
	GOOS=linux go build \
				-o ${APP} \
                -ldflags "-s -w \
				-X  ${PROJECT}/internal/app/config.Release=${RELEASE} \
                -X ${PROJECT}/internal/app/config.Commit=${COMMIT} \
				-X ${PROJECT}/internal/app/config.BuildTime=${BUILD_TIME}" \
				cmd/${APP}/main.go

build: buildlinux

deploy: buildlinux
	./deploy.sh

run:	build
	./${APP} 
test:
	go test -v -race ./...