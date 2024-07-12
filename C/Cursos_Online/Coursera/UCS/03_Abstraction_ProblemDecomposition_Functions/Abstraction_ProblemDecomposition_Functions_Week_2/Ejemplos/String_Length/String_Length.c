/*
* File: String_Basics.c
* Author: Juan Luis Rivera
* Date: 19 de Febrero del 2024
* Description: Ejemplo del uso basico de strings
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

#include <string.h> //Library that will allows us to use methods for strings

int main(int argc, char** argv)
{
    char message [] = "asdawdawcaawfasf";

    int i = 0;
    while (message[i] != '\0') //We obtain the number of characters in the string
    {
        i++;
    }

    //We print the result obtained with the while loop
    printf("Lenght: [%d]\n", i); //We print the size of the string
    
    //We print the string lenght using the strlen function
    printf("Lenght: [%zu]\n", strlen(message));

    //We print the string lenght using the strnlen_s function
    printf("Lenght: [%zu]\n", strnlen(message, sizeof(message))); //This method allows us to avoid the case in which a certaing string doesnt have a null character

    return (EXIT_SUCCESS);
}
