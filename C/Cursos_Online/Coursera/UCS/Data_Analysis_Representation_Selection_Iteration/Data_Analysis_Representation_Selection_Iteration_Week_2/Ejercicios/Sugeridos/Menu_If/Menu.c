/*
*   File: Demostracion de implementacion de un menu simple en C usando if statement
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

#include <ctype.h> //LIbreria utilizada para hacer comprobaciones del tipo de dato char y para convertir de mayuscula a minuscula y visceversa

int main (int argc, char** argv){
    char option;

    printf("******************\n");
    printf("N\tNew Game\n");
    printf("L\tLoad Game\n");
    printf("O\tOptions\n");
    printf("Q\tQuit\n");
    printf("******************\n");
    getchar(); //Getchar is used to avoid problems caused by the enter key being pressed
    scanf("%c", &option); //Scanf introduces problems when using the "\n" character

    option = tolower(option); //Funcion que permite convertir el char a minuscula y asi evitar el error causado por el uso de mayusculas y minusculas

    if (option == 'n') 
    {
        printf("creating new game\n");
    }
    else if (option == 'l')
    {
        printf("Cargando nuevo juego\n");
    }
    else if (option == 'o')
    {
        printf("Mostrando opciones\n");
    }
    else if (option == 'q')
    {
        printf("Saliendo del juego\n");
    }
    else
    {
        printf("Ingresa una opcion valida\n");
    }
    
    return (EXIT_SUCCESS);
}