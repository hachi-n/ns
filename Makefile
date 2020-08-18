go-clean:
	rm -fr ./pkg

go-build:
	go build  -o ./pkg/pigeon ./cli/pigeon

build:
	$(MAKE) go-clean
	$(MAKE) go-build

install:
	$(MAKE) build
	cp -pr ./pkg/pigeon /usr/local/bin

uninstall:
	rm -fr /usr/local/bin/pigeon
