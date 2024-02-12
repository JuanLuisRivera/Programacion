/*
* File: Initialize_Arrays.c
* Author: Juan Luis Rivera
* Date: 28 de enero de 2024
* Description: Ejemplo de inicializacion de arrays
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv)
{
    int a[5]; //A way in which we can initialize arrays without specifying the data in each element

    int b[] = { 1, 2, 3, 4, 5 }; //Another way in which we can initialize arrays specifying the data in each element

    int c[5] = {0}; //Another way in which we can initialize arrays by specifying the data in all elements, in this case all elements will be equal to 0
    
    for (int i = 0; i < 5; i++)
    {
        printf("%d ", a[i]); //This will print the data stored in each element in the memory direction of the array
    }
    printf("\n");
    for (int i = 0; i < 5; i++)
    {
        printf("%d ", b[i]); //This will print each element in the array by the way we initialized them
    }
    printf("\n");
    for (int i = 0; i < 5; i++)
    {
        printf("%d ", c[i]); //This will print all 0 as all elements in the array have been initialized to 0
    }
    printf("\n");

    return (EXIT_SUCCESS);
}
