/*Muestra el funcionamiento de un operador de tres parametros
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    int speed;
    printf("Ingresa velocidad: \n");
    scanf("%d", &speed);
    speed = (speed <= 65)? (65) : (speed <= 70)? (70) : (90);
    switch (speed)
    {
    case 65: 
        printf("\nNo speeding ticket.\n");
        break;
    case 70: 
        printf("\nSpeeding ticket.\n");
        break;
    case 90: 
        printf("\nExpensive speeding ticket.\n");
        break;

    default:
        printf("\nNo valid speed.\n");
    }
    return 0;
}