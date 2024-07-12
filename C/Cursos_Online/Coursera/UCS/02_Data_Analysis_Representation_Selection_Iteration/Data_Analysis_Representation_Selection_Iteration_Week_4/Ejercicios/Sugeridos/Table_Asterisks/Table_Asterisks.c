/*
*   File: Table Asterisks
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*   Description: Program where a table of asterisks is printed in the console specified by the user
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main (int argc, char** argv){
    int width, height;
    printf("Ingresa el ancho de la tabla: ");
    scanf("%d", &width);
    printf("Ingresa el alto de la tabla: ");
    scanf("%d", &height);
    
    for (int i = 0; i < height; i++)
    {
        for (int j = 0; j < width; j++)
        {
            printf("*");
        }
        printf("\n");
    }
    
    
    return (EXIT_SUCCESS);
}