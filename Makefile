
build:
	@./go.sh build

install: build
	@cp ./lico /usr/local/bin
