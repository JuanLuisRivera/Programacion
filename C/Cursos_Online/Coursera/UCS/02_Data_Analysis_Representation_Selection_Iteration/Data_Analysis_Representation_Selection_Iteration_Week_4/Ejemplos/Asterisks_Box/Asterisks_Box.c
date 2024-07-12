/*
*   File: Asterisks Box
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*   Description: Program where we print a box of asterisks depending on the user specification and some size restrictions
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main (int argc, char** argv){
    int width, height;    
    do
    {
        printf("Ingresa el ancho de la caja entre 3 y 20: ");
        scanf("%d", &width);
    } while (width < 3 || width > 20);

    do
    {
        printf("Ingresa el alto de la caja entre 3 y 20: ");
        scanf("%d", &height);
    } while (height < 3 || height > 20);
    
    for (int i = 1; i <= height; i++)
    {
        for (int j = 1; j <= width; j++)
        {
            if (j == 1 || i == 1 || i == height || j == width)
            {
                printf("*");
            }
            else
            {
                printf(" ");
            }
            
        }
        printf("\n");
    }
    
    
    return (EXIT_SUCCESS);
}