clean: all-clean

DATE=$(shell date "+%Y%m%d-%H%M")

all-clean:
	rm -rf bin/*

bin: all-clean
	docker run -t --rm -v ${PWD}:/go/src/github.com/andy-zhangtao/myip -w /go/src/github.com/andy-zhangtao/myip golang:alpine go build -o bin/myip cli/*.go
	docker build -t registry.cn-beijing.aliyuncs.com/vikings/common-tools:myip-$(DATE) .
	docker push registry.cn-beijing.aliyuncs.com/vikings/common-tools:myip-$(DATE)

all: clean bin
