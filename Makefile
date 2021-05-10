build:
	go build -o server main.go

gen:
	go run github.com/99designs/gqlgen generate

run: build
	PORT=80 ./server

serve: build
	PORT=80 ./server

# found this at: https://www.sysleaf.com/go-hot-reload/
watch:
	@ulimit -n 1000 #increase the file watch limit, might required on MacOS
	@reflex -s -r '\.go$$' make run