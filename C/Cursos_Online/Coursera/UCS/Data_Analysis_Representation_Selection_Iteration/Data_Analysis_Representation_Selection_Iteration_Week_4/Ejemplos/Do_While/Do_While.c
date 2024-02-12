/*
*   File: Do While
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*   Description: Program where we validate information entered by the user by using a do while loop
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

int main (int argc, char** argv){
    int score; // Variable to store the user score

    printf("Ingresa tu puntaje: ");
    do
    {
        scanf("%d", &score);
        if (score < 0 || score > 100){
            printf("\n");
            printf("El puntaje debe estar entre 0 y 100\n");
            printf("Ingresa un puntaje valido: ");
        }
    } while (score < 0 || score > 100);   
    
    
    return (EXIT_SUCCESS);
}