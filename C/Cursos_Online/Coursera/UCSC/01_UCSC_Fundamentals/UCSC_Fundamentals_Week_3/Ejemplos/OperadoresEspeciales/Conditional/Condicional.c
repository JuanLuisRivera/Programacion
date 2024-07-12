/*Muestra el funcionamiento de un flujo usando un operador especial
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    int a = 3, b = 2, c;
    c = (a < b)? a: b;
    printf("Valor de c = %d.\n", c);
    return 0;
}