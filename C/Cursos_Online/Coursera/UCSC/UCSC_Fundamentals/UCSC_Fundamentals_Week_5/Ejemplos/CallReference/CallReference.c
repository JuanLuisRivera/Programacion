/*Codigo que muestra el funcionamiento de llamada por referencia en C
 * Juan Luis Rivera
 * 14 de Enero 2024*/

#include <stdio.h>

double average_grades (int how_many, int grades[]){
    int i;
    double sum = 0.0;
    for (i = 0; i < how_many; i++)
    {
        sum += grades[i];
    }

    return (sum / how_many);
}

void print_grades (int how_many, int grades[]){
    int i;
    for (i = 0; i < how_many; i++)
    {
        printf("%d\t", grades[i]);
    }
    
}

int main (void){
    const int SIZE = 5;
    int grades [] = {78, 67, 92, 83, 88};

    print_grades(SIZE, grades);
    printf("\n\n");
    printf("Average = %.2f\n\n", average_grades(SIZE, grades));

    return 0;
}