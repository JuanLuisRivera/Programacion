/*
*   File: Pyramid Asterisks
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*   Description: Program where a pyramid of asterisks is printed in the console specified by the user
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main (int argc, char** argv){
    int height;
    printf("Ingresa el numero maximo de asteriscos: ");
    scanf("%d", &height);
    
    for (int i = 1; i <= height; i++)
    {
        for (int j = 0; j < i; j++)
        {
            printf("*");
        }
        printf("\n");
    }
    
    
    return (EXIT_SUCCESS);
}