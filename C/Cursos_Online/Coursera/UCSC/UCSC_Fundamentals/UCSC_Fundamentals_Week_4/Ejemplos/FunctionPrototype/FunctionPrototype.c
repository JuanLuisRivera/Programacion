/*Codigo que muestra el funcionamiento de una funcion prototipo
 * Juan Luis Rivera
 * 14 de Enero 2024*/

#include <stdio.h>

double square (double); //Prototipo de la funcion square
double cube (double); //Prototipo de la funcion cube

int main (void){
    int how_many = 0, i, j;
    printf("I want square and cube for 1 to n where n is: \n");
    scanf("%d", &how_many);

    printf("\n square and cubes by interval of 0.1.\n");

    for (i = 1; i <= how_many; i++){
        for (j = 0; j < 10; j++)
        {
            printf("\n%lf\t %lf\t %lf\t", i + j /10.0, square (i + j/10.0), cube(i + j/10.0));
        }
    }

    printf("\n\n");
    return 0;
}

double square (double x){ //Definicion de la funcion square
    return (x * x);
}

double cube (double x){ //Definicion de la funcion cube
    return (x * x * x);
}