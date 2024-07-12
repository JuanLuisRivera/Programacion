/*
* File: Complete_Caesar_Cipher.c
* Author: Juan Luis Rivera
* Date: 20 de Febrero del 2024
* Description: Ejemplo de implementacion de cifrado cesar con un cambio introducido por el usuario
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

#include <string.h> //Library that will allows us to use methods for strings
#include <stdbool.h> // Include the standard C boolean definitions
#include <ctype.h> //Allows us to use character oriented methods


int main(int argc, char** argv)
{
    char mensajeOriginal [100];

    int shift = 0;

    bool mensajeValido = false; //Establecemos una bandera para saber si el mensaje es valido

    while (shift == 0) //Verificamos que el usuario introduzca un valor numerico diferente de 0
    {
        printf("Enter shift value: ");

        scanf("%d", &shift);

        if (shift == 0)
        {
            printf("\n");
            printf("Shift quantity cannot be 0\n");
        }

        getchar(); //We use getchar in order to avoid the problem caused by the "\n" character after the user pressed enter
    }

    printf("\nEnter a message: ");
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
        mensajeEncriptado[i] = mensajeOriginal[i] + (shift % 25);
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
        mensajeDesencriptado[i] = mensajeEncriptado[i] - (shift % 25);
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