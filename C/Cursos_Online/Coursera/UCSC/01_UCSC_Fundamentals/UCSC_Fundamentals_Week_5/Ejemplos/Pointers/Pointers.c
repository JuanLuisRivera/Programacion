/*Codigo que muestra el funcionamiento de un apuntador
 * Juan Luis Rivera
 * 14 de Enero 2024*/

#include<stdio.h>

int main (void){
    const int SIZE = 5;

    int grades[] = {78, 67, 53, 49, 84};
    double sum = 0.0;
    double *ptr_to_sum = &sum;
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

    printf("Sum is at %p, or %lu and is %lf.\n", ptr_to_sum, ptr_to_sum, *ptr_to_sum);
    printf("Grades are at %lu to %lu.\n", grades, grades + 5);

    return 0;
}