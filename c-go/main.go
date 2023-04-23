package main

/*
 // C code goes here
#include <stdio.h>

int greet () {
	printf("Hello World! \n");
	return 0;
}

*/
import "C"

func main() {
	C.greet()
}
