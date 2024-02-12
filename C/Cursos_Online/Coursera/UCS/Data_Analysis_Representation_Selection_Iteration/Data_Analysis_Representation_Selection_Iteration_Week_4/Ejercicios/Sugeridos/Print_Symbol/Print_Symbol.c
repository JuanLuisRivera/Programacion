/*
*   File: Print Symbol
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*   Description: Program where the user specifies the number of times an asterisk should be printed
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main (int argc, char** argv){
    int n;
    printf("Ingresa el numero de veces que se imprimira el simbolo: ");
    scanf("%d", &n);
    
    for (int i = 0; i < n; i++)
    {
        printf("*");
    }
    printf("\n");
    
    
    return (EXIT_SUCCESS);
}