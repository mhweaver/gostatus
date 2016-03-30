all:
	go build

run:	all
	./statusbar | lemonbar -f "DejaVu Sans Mono for Powerline"
