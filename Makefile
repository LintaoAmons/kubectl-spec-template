TARGETS = darwin/arm64 linux/amd64 windows/amd64

OUTPUT_DIR = artifact

BINARY_NAME = kubectl-spec-template

all: clean $(TARGETS)

$(TARGETS):
	@echo "Building $@ ..."
	GOOS=$$(echo $@ | cut -d/ -f1) GOARCH=$$(echo $@ | cut -d/ -f2) go build -o $(OUTPUT_DIR)/$(BINARY_NAME)_$$(echo $@ | tr / _) .

clean:
	rm -rf $(OUTPUT_DIR)

build-to-local-bin:
	go build -o ~/.local/bin/$(BINARY_NAME) main.go
	chmod u+x ~/.local/bin/$(BINARY_NAME)

.PHONY: all clean $(TARGETS)
