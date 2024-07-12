/*
 * File:   Assignment_2
 * Author: Juan Luis Rivera
 * Date: 28 de Enero de 2024
 */

#define _CRT_SECURE_NO_WARNINGS
 
#include <stdio.h>
#include <stdlib.h>

#define MAX_LENGTH 100

/*
 * Number Characteristics Programming Assignment
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
		int number = atoi(input);

		// Add your code between this comment
		// and the comment below. You can of
		// course add more space between the
		// comments as needed
		
        if (number % 2 == 0 && number > 0)
        {
            printf("e 1\n");
        }
        else if (number % 2 == 0 && number < 0){
            printf("e -1\n");
        }
        else if (number % 2 != 0 && number > 0)
        {
            printf("o 1\n");
        }
        else if (number % 2 != 0 && number < 0)
        {
            printf("o -1\n");
        }
        else if (number == 0)
        {
            printf("e 0\n");
        }

		// Don't add or modify any code below
		// this comment
		fgets(input, MAX_LENGTH, stdin);
	}

	return 0;
}
