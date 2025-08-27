PROTO_FILES=$(shell find . -name *.proto)

.PHONY: generate
generate:
	protoc --go_out=. --go_opt=paths=source_relative $(PROTO_FILES)

.PHONY: clean
clean:
	rm -f *.pb.go

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  generate - Generate Go code from proto files"
	@echo "  clean    - Remove generated .pb.go files"
	@echo "  help     - Show this help message"