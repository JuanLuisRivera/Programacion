/*
*   File: Nested For Loop example
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main (int argc, char** argv){
    printf("      ");
    
    for (int i = 1; i <= 10; i++)
    {
        printf("%5d\n", i);
    }

    printf("\n");

    //Print a multiplication table

    for (int i = 1; i <= 10; i++)
    {
        printf("%5d", i);
        for (int j = 1; j <= 10; j++)
        {
            printf("%5d", i * j);
        }
        printf("\n");
    }
    

    
    
    return (EXIT_SUCCESS);
}