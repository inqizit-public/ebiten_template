
CWD := $(shell pwd)
DIST := ${CWD}/dist
RESOURCES_PATH := ${CWD}/resources
DIST_JS := ${DIST}/js
GOROOT := $(shell go env GOROOT)

run: 
	go run cmd/main.go
.PHONY: run

web: 
	# go run github.com/hajimehoshi/wasmserve@latest ./cmd/main.go
	# go install github.com/hajimehoshi/wasmserve@latest 
	echo "available at http://localhost:8080/"
	wasmserve ./cmd/main.go
.PHONY: web

dist-js:
	env GOOS=js GOARCH=wasm go build -o ./dist/js/ebiten_template.wasm ./cmd/main.go
	-mkdir -p dist/js
	cp ${GOROOT}/misc/wasm/wasm_exec.js ${DIST_JS}
	cp ${RESOURCES_PATH}/deploy/js/index.html ${DIST_JS}
.PHONY: dist-js

deploy-js: dist-js
	-rm -rf _tmp
	git clone https://github.com/inqizit-public/ebiten_template.git _tmp
	cd _tmp; git checkout gh-pages; rm -rf *; cp ${DIST_JS}/* .; git add .; git commit -m "deploy"; git push
	-rm -rf _tmp
.PHONY: deploy-js

build:
	go build -o ./dist/desktop/ebiten_template ./cmd/main.go
.PHONY: build