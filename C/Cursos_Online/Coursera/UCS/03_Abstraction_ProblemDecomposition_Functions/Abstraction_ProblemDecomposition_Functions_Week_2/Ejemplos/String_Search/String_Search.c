/*
* File: String_Search.c
* Author: Juan Luis Rivera
* Date: 19 de Febrero del 2024
* Description: Ejemplo de implementacion de busqueda en strings
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

#include <string.h> //Library that will allows us to use methods for strings

int main(int argc, char** argv)
{
    char data [] = "asdawdawcaawfasbf"; //Proponemos un string para probar la busqueda

    int i = 0; //We create a variable to count the number of characters until we find the character we are looking for

    int lenght = strnlen(data, sizeof(data)); //We obtain the number of characters in the string

    while (data[i] != 'b' && i < lenght) //We run the loop until we find the character 'b' or we reach the end of the string
    {
        i++; //we increment the counter to check the next character
    }

    if (i > lenght)
    {
        printf("Character not found\n"); //Case in which we traveled the entire string without finding the character
    }
    else{
        printf("Character: [%c]\n", data[i]); //We print the character
        printf("Character location: [%d]\n", i); //We print the location of the character
    }
    
    //We will find the given character using a standard library function
    char* index = NULL; //We create a pointer to store the memory adress of the character and we initialize it to NULL
    index = strchr(data, 'b'); //We use the strchr function to find the character 'b' in the string

    if (index != NULL) //Case in which we found the character and so the function returns a memory adress, aka, it is not NULL
    {
        printf("Character adress: [%p]\n", index); //We print the character adress using the proper format specifier "%p"
        printf("Data start location: [%p]\n", &data[0]); //We print the location of the first character in the string

        char* inicioDato = &data[0]; //We create a pointer to store the location of the first character in the string
        printf("Character location: [%ld]\n", index - inicioDato); //We print the location of the character using the proper format specifier "%ld"
    }
    else
    {
        printf("Character not found\n"); //Case in which we obtained a NULL pointer
    }
    

    return (EXIT_SUCCESS);
}