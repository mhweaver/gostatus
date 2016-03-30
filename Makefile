all:
	go build

run:	all
	./statusbar | lemonbar
