default: all

load: load.c
	gcc load.c -o load -g -ggdb -O0 -std=c11 -Wall -Wextra

run-load:
	./load

all: load run-load

clean:
	rm ./load
