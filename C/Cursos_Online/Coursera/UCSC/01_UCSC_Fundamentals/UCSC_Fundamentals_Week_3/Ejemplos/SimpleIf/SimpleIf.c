/*Muestra el funcionamiento de un flujo de if simple
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    int speed;
    printf("Enter your speed as an integer: \n");
    scanf("%d", &speed);
    if (speed < 65)
        printf ("\nNo speed ticket.\n");
    else
        printf("\nSpeeding ticket.\n");
    return 0;
}