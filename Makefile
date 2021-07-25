build:
	go build -o puzzled main.go

run: build
	./puzzled

watch:
	ulimit -n 1000 #increase the file watch limit, might required on MacOS
	reflex -s -r '\.go$$' make run
