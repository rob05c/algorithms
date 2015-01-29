CC=clang++
OPTIONS=-O3 -std=c++11 -g -Wfatal-errors
LINK_OPTIONS= #-ltbb
all: mergesort
mergesort: main.o
	$(CC) $(OPTIONS) main.o -o mergesort $(LINK_OPTIONS)
main.o: 
	$(CC) $(OPTIONS) -c main.cpp -o main.o
clean:
	rm -f *.o mergesort
