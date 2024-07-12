/*
*   File: User Print
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*   Description: Program where the user specifies the maximum number and it print the odd number until that maximum number
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main (int argc, char** argv){
    int n;
    printf("Ingresa un numero maximo: ");
    scanf("%d", &n);
    
    for (int i = 0; i <= n; i++)
    {
        if (i % 2 != 0)
        {
            printf("%d\n", i);
        }
    }
    
    
    return (EXIT_SUCCESS);
}