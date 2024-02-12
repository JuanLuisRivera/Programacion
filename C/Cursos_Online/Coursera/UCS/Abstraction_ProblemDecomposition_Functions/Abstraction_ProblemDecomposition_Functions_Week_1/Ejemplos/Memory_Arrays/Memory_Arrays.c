/*
* File: Memory_Arrays.c
* Author: Juan Luis Rivera
* Date: 28 de enero de 2024
* Description: Ejemplo del funcionamiento de arrays en memoria
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv)
{
    int scores[5] = {100 , 90, 80, 70, 60};

    for (int i = 0; i < 5; i++) //We print the array and the memory address of each element
    {
        printf("Scores: [%d] Address: %p Contents: %d\n", i, &scores[i], scores[i]); //We use the format specifier "%p" to print adresses
    }
    
    return (EXIT_SUCCESS);
}
