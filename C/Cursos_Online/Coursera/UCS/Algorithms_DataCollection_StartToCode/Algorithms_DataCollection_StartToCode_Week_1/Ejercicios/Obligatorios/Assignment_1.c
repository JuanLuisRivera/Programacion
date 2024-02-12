/*Codigo que imprime los 10 primeros elementos de la tabla periodica
 * Juan Luis Rivera
 * 19 de Enero 2024*/

#define _CRT_SECURE_NO_WARNINGS
 
#include <stdio.h>
#include <stdlib.h>

#define MAX_LENGTH 100

/*
 * Programming Assignment 1
 */
int main(int argc, char** argv)
{
	// IMPORTANT: Only add code in the section
	// indicated below. The code I've provided
	// makes your solution work with the 
	// automated grader on Coursera
	char input[MAX_LENGTH];
	fgets(input, MAX_LENGTH, stdin);
	while (input[0] != 'q')
	{
		// Add your code between this comment
		// and the comment below. You can of
		// course add more space between the
		// comments as needed

        //Prints the symbol of the first element
        printf("H\n");

        //Prints the symbol of the second element
        printf("He\n");

        //Prints the symbol of the thrid element
        printf("Li\n");

        //Prints the symbol of the fourth element
        printf("Be\n");

        //Prints the symbol of the fifth element
        printf("B\n");

        //Prints the symbol of the sixth element
        printf("C\n");

        //Prints the symbol of the seventh element
        printf("N\n");

        //Prints the symbol of the eighth element
        printf("O\n");

        //Prints the symbol of the ninth element
        printf("F\n");

        //Prints the symbol of the tenth element
        printf("Ne\n");
		
		// Don't add or modify any code below
		// this comment
		fgets(input, MAX_LENGTH, stdin);
	}

	return 0;
}