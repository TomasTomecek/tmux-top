default: all

load: load.c
	gcc load.c -o load -std=c11 -Wall -Wextra

run-load:
	./load

all: load run-load

