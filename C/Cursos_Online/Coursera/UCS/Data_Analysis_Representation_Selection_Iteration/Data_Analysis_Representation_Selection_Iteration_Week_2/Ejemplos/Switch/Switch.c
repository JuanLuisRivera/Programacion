/*
*   File: Ejemplo de uso de switch 
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h> 

#include <ctype.h> //LIbreria utilizada para hacer comprobaciones del tipo de dato char y para convertir de mayuscula a minuscula y visceversa

int main (int argc, char** argv){
    char option;

    printf("Pizza? (y/n)\n");

    scanf("%c", &option);

    switch (option)
    {
    case 'y':
        printf ("Toma pizza :).\n");
        break;
    
    
    case 'n':
        printf("No pizza then :C. \n");
        break;

    default:
        printf("Esa no es una opcion >:B");
        break;
    }
    return (EXIT_SUCCESS);
}