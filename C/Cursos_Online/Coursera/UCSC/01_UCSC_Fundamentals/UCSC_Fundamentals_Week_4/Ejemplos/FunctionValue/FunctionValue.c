/*Codigo que muestra el funcionamiento de una funcion llamada por valor
 * Juan Luis Rivera
 * 14 de Enero 2024*/

#include <stdio.h>

int funcionValor (int);

int main (void){
    int i = 5;
    printf("Valor de funcionValor(i) = %d.\n", funcionValor(i));

    printf("Valor de i = %d.\n", i);

    return 0;
}

int funcionValor(int n){
    n = n + 1;
    return n;
}