/*Programa que calcula la distancia entre 2 puntos en dos dimensiones
 * Juan Luis Rivera
 * 22 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>


int main (int argc, char** argv){
    //We define a variable of type float where to store the data from the user
    float pointOne_x;
    float pointOne_y; 

    float pointTwo_x;
    float pointTwo_y;

    //Prompts for requesting the user data
    printf("Enter x for first point: \n");
    scanf("%f", &pointOne_x);

    printf("Enter y for first point: \n");
    scanf("%f", &pointOne_y);

    printf("Enter x for second point: \n");
    scanf("%f", &pointTwo_x);

    printf("Enter y for second point: \n");
    scanf("%f", &pointTwo_y);

    printf("\n");

    float distancia = sqrtf(powf((pointOne_x - pointTwo_x), 2) + powf((pointOne_y - pointTwo_y), 2));

    printf("Distancia entre los puntos = %f\n", distancia);

    return (EXIT_SUCCESS);
}