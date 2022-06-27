CMD=go
BIN_PATH=bin
SRC_PATH=src

all: clean build install

build:
	$(CMD) build -o $(BIN_PATH)/qocker $(SRC_PATH)/*
install:
	cp bin/qocker /usr/bin/qocker
	cp bin/qocker /usr/local/bin/qocker

uninstall:
	rm -rf /usr/bin/qocker /usr/local/bin/qocker
	rm -rf /bin/qocker

clean: uninstall

