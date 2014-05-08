stylus:
	stylus ./public/css/style.styl
bindata:
	go-bindata -o ./templates.go templates
build:
	go build ./...
