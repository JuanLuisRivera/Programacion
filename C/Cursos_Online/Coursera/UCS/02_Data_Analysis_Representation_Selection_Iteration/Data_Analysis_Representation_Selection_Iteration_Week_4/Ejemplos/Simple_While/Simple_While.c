/*
*   File: Simple While
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*   Description: Program where we validate information entered by the user by using a simple while loop
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main (int argc, char** argv){
    int score;
    printf("Ingresa tu puntaje: ");
    scanf("%d", &score);
    
    while (score < 0 || score > 100)
    {
        printf("\n");
        printf("El puntaje debe estar entre 0 y 100\n");
        printf("Ingresa un puntaje valido: ");
        scanf("%d", &score);
    }
    
    return (EXIT_SUCCESS);
}