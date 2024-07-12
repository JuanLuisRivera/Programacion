/*
* File: String_Basics.c
* Author: Juan Luis Rivera
* Date: 19 de Febrero del 2024
* Description: Ejemplo del uso basico de strings
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv)
{
    char message [] = "FTATRIST";

    int i = 0;
    while (message[i] != '\0') //We print the number data, the adress and the content of each element in the array
    {
        printf("Mensaje: [%d] \t Direccion: [%p] \t Contenido: [%c]\n", i, &message[i], message[i]);
        i++;
    }

    printf("Message: [%s]\n", message); //We print the entire string with the appropiate format specifier "%s"
    

    return (EXIT_SUCCESS);
}
