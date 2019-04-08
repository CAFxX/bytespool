
.PHONY: clean
clean:
	rm *.a *.s *.go

.PHONY: generate
generate: clean
	go run gen/main.go | gofmt > bytespool.go

.PHONY: build
build: generate
	go build -o pool.a .

.PHONY: disasm
disasm: build
	go tool objdump pool.a > pool.s

.PHONY: test
test: generate
	go test .

.PHONY: all
all: test disasm
	# done