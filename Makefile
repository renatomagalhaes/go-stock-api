build:
	GOOS=linux GOARCH=amd64 go build -o app cmd/app/main.go

run:
	./app

clean:
	rm app