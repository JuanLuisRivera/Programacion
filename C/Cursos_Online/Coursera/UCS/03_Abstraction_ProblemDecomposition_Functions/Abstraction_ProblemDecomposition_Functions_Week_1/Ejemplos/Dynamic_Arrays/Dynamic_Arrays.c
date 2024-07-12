/*
* File: Dynamic_Arrays.c
* Author: Juan Luis Rivera
* Date: 28 de enero de 2024
* Description: Ejemplo de creacion de arrays en tiempo de ejecucion
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv)
{
    int n;

    scanf("%d", &n);

    int* scores = malloc(n * sizeof(int)); //We create an array in memory and we specify the number of elements that will be stored in the array
    /*
    * We use "int*" in order to specify that we will be creating a pointer to an integer in this case the "scores" variable
    * As the pointer is pointing to the "scores" variable, we will be using its memory address as the base location for our array
    * We will be using the malloc function in order to create the array in memory that is, we will be be allocating memory by ourselves
    * We will use the "sizeof" function in order to specify the number of bytes that will be stored by each element and the "n" variable to specify the number of elements that will be stored in the array
    */

    printf("Total array size: %lu\n", n * sizeof(int)); //We use the format specifier "%lu" to print the total number of bytes that will be stored in the array
    printf("Array elements size: %lu\n", sizeof(scores[0])); //We use the format specifier "%lu" to print the size of each element in the array
    printf("Number of elements: %lu\n", (n * sizeof(int)) / sizeof(scores[0])); //We print the number of elements that will be stored in the array
    printf("\n");

    for (int i = 0; i < n; i++) //We insert data into the array
    {
        printf("Enter score %d: ", i + 1);
        scanf("%d", &scores[i]);
        while (scores[i] < 0 || scores[i] > 100)
        {
            printf("\n");
            printf("Invalid score. Enter score %d: ", i + 1);
            scanf("%d", &scores[i]);
        }
    }

    printf("\n");
    for (int i = 0; i < n; i++) //We print the array
    {
        printf("%d ", scores[i]);
    }
    printf("\n");

    free(scores); //We free the memory allocated to the array
    scores = NULL; //This way the memory wont be used as the pointer will be freed as well

    return (EXIT_SUCCESS);
}
