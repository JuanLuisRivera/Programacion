/*Muestra el funcionamiento de operadores logicos y evaluacion en corto circuito
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    int outside, weather;
    printf("Enter if outside 1 true 0 false: \n");
    scanf("%d", &outside);
    printf("Enter if rain 1 true 0 false: \n");
    scanf("%d", &weather);
    if (outside && weather)
        printf ("\nUse umbrella.\n");
    else
        printf("\nDress normally.\n");
    return 0;
}