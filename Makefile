
.PHONY: clean
clean:
	rm -f *.a *.s *.go

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

.PHONY: bench
bench: test
	bash -c "benchstat <(go test -run=^$$ -bench=. -count=10 -benchtime=0.1s -benchmem ./test)"

.PHONY: all
all: test bench disasm
	# done