ifeq ($(OS),Windows_NT)
    # If the OS is Windows, use Windows commands
    RM = del /Q
    MKDIR = mkdir
    TARGET = bin\main.exe
		AIR_CONFIG = .air.windows.toml
else
    # If the OS is not Windows, assume it's a Unix-like system (Linux, macOS)
    RM = rm -f
    MKDIR = mkdir -p
    TARGET = bin/main
		AIR_CONFIG = .air.mac.toml
endif

build:
	go build -o $(TARGET) -v
run:
	go run main.go

dev:
	air -c $(AIR_CONFIG)

test:
	go test -v ./...
clean:
	go clean
	$(RM) -f $(TARGET)