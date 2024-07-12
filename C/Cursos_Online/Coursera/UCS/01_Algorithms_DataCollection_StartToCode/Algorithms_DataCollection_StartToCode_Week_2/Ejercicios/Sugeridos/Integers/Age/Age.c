/*Ejemplo de impresion con especificadores de enteros
 * Juan Luis Rivera
 * 19 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>


int main (int argc, char** argv){
    //We define the int variable
    int age = 25; 

    //We print the age using the proper format specifier for integers "%d"
    printf("Age: %d\n", age);

    return (EXIT_SUCCESS);
}