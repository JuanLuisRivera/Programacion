/*Muestra el tamaño de los tipos de dato fundamentales
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    printf("El tamaño de los tipos de dato en mi sistema es: \n");
    printf("Int es igual a: %lu bytes.\n", sizeof(int));
    printf("Long es igual a: %lu bytes.\n", sizeof(long));
    printf("Char es igual a: %lu bytes.\n", sizeof(char));
    printf("Float es igual a: %lu bytes.\n", sizeof(float));
    printf("Double es igual a: %lu bytes.\n", sizeof(double));
    printf("Long double es igual a: %lu bytes.\n", sizeof(long double));
    return 0;
}