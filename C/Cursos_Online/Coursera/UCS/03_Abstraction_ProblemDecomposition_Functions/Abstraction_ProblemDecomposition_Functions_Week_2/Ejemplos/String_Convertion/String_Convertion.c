/*
* File: String_Convert.c
* Author: Juan Luis Rivera
* Date: 19 de Febrero del 2024
* Description: Ejemplo de implementacion de la conversion de datos a partir de un string
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

#include <string.h> //Library that will allows us to use methods for strings

int main(int argc, char** argv)
{
    //Convert the data from string to float
    char data [100]; //We create a variable to store the string
    float floatData;

    printf("Enter a float: "); //We print a message to the user
    fgets(data, sizeof(data), stdin); //We get the string from the user using the "fgets" function which reads until the enter key is pressed
    //The fgets function tackes as arguments the variable where the string will be stored, the max size of the string and the input stream and returns the string with a null terminator

    printf("You entered: %s\n", data); //We print the number inputted
    
    //Convert the data from string to float
    floatData = atof(data); //We use the stdlib atof function to convert the string to a float
    printf("Float data: %f\n", floatData); //We print the float data


    //Convert the data from string to integer
    char intString [100]; //We create a variable to store the numeric string
    int intData; //We create a variable to store the integer data
    printf("\n\nEnter an integer: "); //We request the user to enter an integer
    fgets(intString, sizeof(intString), stdin); //We get the string from the user using the "fgets" function

    printf("Integer string data: %s\n", intString); //We print the string containing the integer data

    intData = atoi(intString); //We use the stdilib atoi function to convert the string to an integer
    printf("Integer data: %d\n", intData); //We print the integer data

    return (EXIT_SUCCESS);
}