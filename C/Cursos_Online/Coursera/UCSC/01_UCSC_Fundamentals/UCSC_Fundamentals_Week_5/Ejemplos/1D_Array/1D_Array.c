/*Codigo que muestra el funcionamiento basico de un arreglo
 * Juan Luis Rivera
 * 14 de Enero 2024*/

#include<stdio.h>

int main (void){
    const int SIZE = 5;

    int grades[] = {78, 67, 53, 49, 84};
    double sum = 0.0;
    int i;

    printf("Grades are: \n");

    for (i = 0; i < SIZE; i++)
    {
        printf("%d\t", grades[i]);
    }
    
    printf("\n\n");

    for (i = 0; i < SIZE; i++)
    {
        sum = sum + grades[i];
    }

    printf("Average: %.2f\n\n", sum / SIZE);

    return 0;
}