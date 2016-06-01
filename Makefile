all:
	go build

run:	all
	./gostatus | lemonbar -f "DejaVu Sans Mono for Powerline"

gostatus:	all

install:	gostatus
	cp $< ${HOME}/bin
