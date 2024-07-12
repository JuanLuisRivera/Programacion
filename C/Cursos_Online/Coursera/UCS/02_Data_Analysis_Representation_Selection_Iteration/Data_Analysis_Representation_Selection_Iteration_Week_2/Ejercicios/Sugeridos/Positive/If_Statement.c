/*
*   File: Ejercicio 1 de If
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h> 

#include <ctype.h> //LIbreria utilizada para hacer comprobaciones del tipo de dato char y para convertir de mayuscula a minuscula y visceversa

int main (int argc, char** argv){
    int number;

    printf("Numero Positivo: ");
    scanf("%d", &number);

    if (number >= 0) 
    {
        printf("Positivo.\n");
    }
    else if (number < 0)
    {
        printf("Negativo.\n");
    }
    
    return (EXIT_SUCCESS);
}