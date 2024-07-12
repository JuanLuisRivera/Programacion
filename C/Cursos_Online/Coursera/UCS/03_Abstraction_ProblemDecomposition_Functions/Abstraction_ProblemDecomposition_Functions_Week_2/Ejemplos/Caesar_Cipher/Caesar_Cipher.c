/*
* File: Caesar_Cipher.c
* Author: Juan Luis Rivera
* Date: 19 de Febrero del 2024
* Description: Ejemplo de implementacion de cifrado cesar con un cambio de 5 en el alfabeto
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

#include <string.h> //Library that will allows us to use methods for strings
#include <stdbool.h> // Include the standard C boolean definitions
#include <ctype.h> //Allows us to use character oriented methods

#define SHIFT_AMOUNT 5

int main(int argc, char** argv)
{
    char mensajeOriginal [100]; //Creamos un arreglo para guardar el mensaje introducido por el usuario

    bool mensajeValido = false; //Establecemos una bandera para saber si el mensaje es valido

    printf("Enter a message: ");
    fgets(mensajeOriginal, sizeof(mensajeOriginal), stdin); //Obtenemos la cadena de caracteres escrita por el usuario
    int mensajeOriginalSize = strnlen(mensajeOriginal, sizeof(mensajeOriginal)) - 1; //Creamos una variable para obtener la longitud del mensaje y evitamos la inclusion del caracter de salto de linea al restar 1 del tamaño obtenido por "sizof"

    while (!mensajeValido)
    {
        mensajeValido = true;

        for (int i = 0; i < mensajeOriginalSize; i++)
        {
            if (!isupper(mensajeOriginal[i])) //Verificamos que todos los caracteres sean mayusculas
            {
                printf("\n");
                printf("Enter a message with only uppercase letters: ");
                fgets(mensajeOriginal, sizeof(mensajeOriginal), stdin);
                mensajeOriginalSize = strnlen(mensajeOriginal, sizeof(mensajeOriginal) - 1);
                mensajeValido = false;
                break;
            }
        }
    }


    char* mensajeEncriptado = malloc((mensajeOriginalSize + 1) * sizeof(char));
    for (int i = 0; i < mensajeOriginalSize; i++)
    {
        mensajeEncriptado[i] = mensajeOriginal[i] + SHIFT_AMOUNT;
        if (mensajeEncriptado[i] > 'Z')
        {
            mensajeEncriptado[i] -= 'Z' - 'A' + 1;
        }
    }
    mensajeEncriptado[mensajeOriginalSize] = '\0';


    int mensajeEncriptadoSize = mensajeOriginalSize;
    char* mensajeDesencriptado = malloc((mensajeEncriptadoSize + 1) * sizeof(char));
    for (int i = 0; i < mensajeEncriptadoSize; i++)
    {
        mensajeDesencriptado[i] = mensajeEncriptado[i] - SHIFT_AMOUNT;
        if (mensajeDesencriptado[i] < 'A')
        {
            mensajeDesencriptado[i] += 'Z' - 'A' + 1;
        }
    }

    mensajeDesencriptado[mensajeOriginalSize] = '\0';
    printf("Original message: %s\n", mensajeOriginal);
    printf("Encrypted message: %s\n", mensajeEncriptado);
    printf("Decrypted message: %s\n", mensajeDesencriptado);

    free(mensajeEncriptado);
    mensajeEncriptado = NULL;
    free(mensajeDesencriptado);
    mensajeDesencriptado = NULL;

    return (EXIT_SUCCESS);
}