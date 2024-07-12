/*
*   File: Simple If statement example
*   Author: Juan Luis Rivera
*   Date: 28 de Enero de 2024
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

#include <ctype.h> //LIbreria utilizada para hacer comprobaciones del tipo de dato char y para convertir de mayuscula a minuscula y visceversa

int main (int argc, char** argv){
    char answer;

    printf("Pizza? (y, n)");
    scanf("%c", &answer);
    
    answer = tolower(answer); //Funcion que permite convertir el char a minuscula

    if (answer == 'y' || answer == 'Y') 
    {
        printf("Toma pizza :)\n");
    }
    else if (answer == 'n') //Ya no es necesaria la posibilidad de que el caracter 'n' sea en mayuscula porque se convierte a minuscula con la funcion tolower()
    {
        printf("No hay pizza :(\n");
    }
    else{
        printf("Esa no es una opcion >:B\n");
    }
    
    return (EXIT_SUCCESS);
}