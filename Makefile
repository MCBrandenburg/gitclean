build-all: clean-dist build-darwin build-freebsd build-linux build-windows

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./dist/gitclean
	tar -C dist/ -czf ./dist/gitclean.darwin.amd64.tar.gz gitclean
	rm -rf ./dist/gitclean

	GOOS=darwin GOARCH=386 go build -o ./dist/gitclean
	tar -C dist/ -czf ./dist/gitclean.darwin.386.tar.gz gitclean
	rm -rf ./dist/gitclean

build-freebsd:
	GOOS=freebsd GOARCH=amd64 go build -o ./dist/gitclean
	tar -C dist/ -czf ./dist/gitclean.freebsd.amd64.tar.gz gitclean
	rm -rf ./dist/gitclean

	GOOS=freebsd GOARCH=386 go build -o ./dist/gitclean
	tar -C dist/ -czf ./dist/gitclean.freebsd.386.tar.gz gitclean
	rm -rf ./dist/gitclean

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./dist/gitclean
	tar -C dist/ -czf ./dist/gitclean.linux.amd64.tar.gz gitclean
	rm -rf ./dist/gitclean

	GOOS=linux GOARCH=386 go build -o ./dist/gitclean
	tar -C dist/ -czf ./dist/gitclean.linux.386.tar.gz gitclean
	rm -rf ./dist/gitclean

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./dist/gitclean.exe
	tar -C dist/ -czf ./dist/gitclean.windows.amd64.tar.gz gitclean.exe
	rm -rf ./dist/gitclean.exe

	GOOS=windows GOARCH=386 go build -o ./dist/gitclean.exe
	tar -C dist/ -czf ./dist/gitclean.windows.386.tar.gz gitclean.exe
	rm -rf ./dist/gitclean.exe

clean-dist:
	rm -rf dist/*