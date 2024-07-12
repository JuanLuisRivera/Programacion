/*
* File: CSV.c
* Author: Juan Luis Rivera
* Date: 19 de Febrero del 2024
* Description: Ejemplo de procesamiento de archivos CSV usando los metodos vistos
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

#include <string.h> //Library that will allows us to use methods for strings

int main(int argc, char** argv)
{
    struct CSV
    {
        float x;
        float y;
    };

    typedef struct CSV CSV;
    CSV point;
    
    char pointString [100];
    printf("Enter x and y for point: ");
    fgets(pointString, sizeof(pointString), stdin);

    int commaIndex = -1;
    char* result = NULL;
    while (commaIndex == -1)
    {
        result = strchr(pointString, ',');
        if (result != NULL)
        {
            char* pointStart = &pointString[0];
            commaIndex = result - pointStart;
        }

        else
        {
            printf("\n");
            printf("Error: No comma found in the string\n");
            printf("Enter x and y for point: ");
            fgets(pointString, sizeof(pointString), stdin);
        }
        
    }
    
    //X part
    char* xString = malloc((commaIndex + 1) * sizeof(char)); //Se declara un apuntador con el objetivo de resevar suficiente memoria para la variable "x" del tipo de dato "Vector"
    strncpy(xString, pointString, commaIndex); //Se copia la cadena de caracteres en la direccion de memoria a la que apunta "xString"
    xString[commaIndex] = '\0'; //Se agrega el caracter de fin de cadena al final de la direccion de memoria a la que apunta "xString"
    point.x = atof(xString); //Se convierte el string a float usando la funcion atof

    free(xString); //Se liberan los recursos reservados en el sistema
    xString = NULL; //Se evita el posterior uso de la memoria haciendo que el apuntador apunte a NULL

    //Y part
    int length = strnlen(pointString, sizeof(pointString)) - 1; //Se calcula la longitud de la cadena de caracteres sin incluir el caracter de nueva cadena "\n"

    char* yString = malloc((length - commaIndex) * sizeof(char)); //Se declara un apuntador con el objetivo de resevar suficiente memoria para la variable "y" del tipo de dato "Vector"
    int offset = commaIndex + 1; //Se calcula un "offset" con el objetivo de iniciar la adquision de datos en el siguiente elemento a partir de la coma
    for (int i = 0; i < length - commaIndex - 1; i++) //Se verifica que no se tome el caracter de salto de linea en el nuevo string
    {
        yString[i] = pointString[offset + i]; //Se copian todos los caracteres de la cadena en el espacio reservado anteriormente por el apuntador
    }

    yString[length - commaIndex - 1] = '\0'; //Se agrega el caracter de fin de cadena al final de la direccion de memoria a la que apunta "yString" para tener una string bien formada
    point.y = atof(yString); //Se convierte el string a float usando la funcion atof
    
    free(yString); //Se libera la memoria reservada en el sistema
    yString = NULL; //Se evita el posterior uso de la memoria haciendo que el apuntador apunte a NULL

    //Se imprime la informacion del punto
    printf("\n");
    printf("X: %f\n", point.x);
    printf("Y: %f\n", point.y);
    

    return (EXIT_SUCCESS);
}