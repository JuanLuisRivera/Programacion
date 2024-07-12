/*
* File: String_Input.c
* Author: Juan Luis Rivera
* Date: 19 de Febrero del 2024
* Description: Ejemplo de implementacion de user input using strings
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

#include <string.h> //Library that will allows us to use methods for strings

int main(int argc, char** argv)
{
    char data [100]; //We create a variable to store the string

    printf("Enter a string: "); //We print a message to the user
    fgets(data, sizeof(data), stdin); //We get the string from the user using the fgets function which reads until the enter key is pressed
    //The fgets function tackes as arguments the variable where the string will be stored, the max size of the string and the input stream and returns the string with a null terminator

    printf("You entered: %s\n", data); //We print the string inputted
    

    return (EXIT_SUCCESS);
}