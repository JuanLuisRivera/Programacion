/*Convierte fahrenheit a celcius
 * Juan Luis Rivera
 * 8 de Enero 2024*/

#include <stdio.h>

int main(void){
    int fahrenheit, celsius;
    printf("Enter Fahrenheit as an integer: \n");
    scanf("%d", &fahrenheit);
    celsius = (fahrenheit - 32)/1.8; //Se lleva acabo la conversion
    printf("\n %d fahrenheit is %d celsius.\n", fahrenheit, celsius);
    return 0;
}
