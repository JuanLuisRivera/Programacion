/*
*   File: Simple For Loop example
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main (int argc, char** argv){
    int n;
    printf("Ingresa un numero de cuadrados: ");
    scanf("%d", &n);
    
    for (int i = 0; i <= n; i++)
    {
        printf("Cuadrado de %d es %d\n", i, i*i);
    }
    
    
    return (EXIT_SUCCESS);
}