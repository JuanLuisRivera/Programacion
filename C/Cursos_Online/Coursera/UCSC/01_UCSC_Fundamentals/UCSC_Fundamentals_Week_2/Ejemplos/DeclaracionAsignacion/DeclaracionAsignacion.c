/*Muestra la forma en que se declaran e inicializan tipos y variables
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    int a = 5, b = 7, c = 6; //Declara e Inicializa
    double average = 0.0; 

    printf("a = %d, b = %d, c = %d\n", a, b, c);
    average = (a + b + c) / 3.0;
    printf("average = %lf\n", average);
    return 0;
}