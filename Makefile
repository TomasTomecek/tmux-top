default:
	gcc main.c -o main -std=c11 -Wall -Wextra

run:
	./main

all: default run

