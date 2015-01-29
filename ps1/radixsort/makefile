CC=clang++
OPTIONS=-O3 -std=c++11 -g -Wfatal-errors
LINK_OPTIONS= #-ltbb
all: radixsort
radixsort: main.o
	$(CC) $(OPTIONS) main.o -o radixsort $(LINK_OPTIONS)
main.o: 
	$(CC) $(OPTIONS) -c main.cpp -o main.o
clean:
	rm -f *.o radixsort
