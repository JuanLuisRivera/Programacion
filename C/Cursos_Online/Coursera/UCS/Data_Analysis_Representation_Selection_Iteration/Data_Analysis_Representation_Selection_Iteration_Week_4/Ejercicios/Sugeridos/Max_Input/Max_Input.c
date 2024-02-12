/*
*   File: Max Input
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*   Description: Program where the maximum number entered by the user is stored until the user enters a "-1"
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

int main (int argc, char** argv){
    int number = INT_MIN; // Variable to store the maximum number
    int aux = 0; // Auxiliar variable to store the numbers entered by the user
    int i = 0; // Counter that counts the numbers entered by the user
    printf("Ingresa los numeros que necesites y cuando termines ingresa un -1: ");

    while (aux != -1)
    {
        scanf("%d", &aux);
        if (i == 0 && aux == -1){
            printf("Ningun numero ingresado.\nSaliendo del programa\n");
            return (EXIT_SUCCESS);
        }

        else{
            if (aux > number)
            {
                number = aux;
                i ++;
            }
        }
    }

    printf("El numero maximo es: %d\n", number);
    
    
    
    return (EXIT_SUCCESS);
}