/*
* File: Simple_Arrays.c
* Author: Juan Luis Rivera
* Date: 28 de enero de 2024
* Description: Ejemplo de arrays simples
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv)
{
    int scores[5];

    for (int i = 0; i < 5; i++) //We insert data into the array
    {
        printf("Enter score %d: ", i + 1);
        scanf("%d", &scores[i]);
        while (scores[i] < 0 || scores[i] > 100)
        {
            printf("\n");
            printf("Invalid score. Enter score %d: ", i + 1);
            scanf("%d", &scores[i]);
        }
    }

    for (int i = 0; i < 5; i++) //We print the array
    {
        printf("Score %d: %d\n", i + 1, scores[i]);
    }
    
    return (EXIT_SUCCESS);
}
