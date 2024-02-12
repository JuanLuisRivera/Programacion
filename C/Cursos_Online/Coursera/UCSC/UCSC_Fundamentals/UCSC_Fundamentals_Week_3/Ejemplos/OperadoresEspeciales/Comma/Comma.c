/*Muestra el funcionamiento de un flujo usando la coma como operador especial
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    int a, b, c;
    c = (a = 0, b = 1);
    printf("Valor de c = %d.\n", c);
    return 0;
}