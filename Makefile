go-clean:
	rm -fr ./pkg

go-build:
	go build  -o ./pkg/pigeon ./cli/pigeon

go-generate:
	go generate  ./cli/pigeon

build:
	$(MAKE) go-clean
	$(MAKE) go-generate
	$(MAKE) go-build

install:
	$(MAKE) build
	cp -pr ./pkg/pigeon /usr/local/bin

uninstall:
	rm -fr /usr/local/bin/pigeon
