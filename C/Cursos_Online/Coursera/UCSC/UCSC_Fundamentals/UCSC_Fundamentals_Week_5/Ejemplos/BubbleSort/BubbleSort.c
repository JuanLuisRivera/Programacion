/*Codigo que muestra el funcionamiento del algoritmo bubble sort en C
 * Juan Luis Rivera
 * 15 de Enero 2024*/

#include <stdio.h>

void swap(int *a, int *b){
    int temp = *a;
    *a = *b;
    *b = temp;
}

void print_array (int how_many, int data [], char *str){
    int i;
    printf("%s", str);

    for (i = 0; i < how_many; i++)
    {
        printf("%d\t", data[i]);
    }
}

void bubble(int data[], int how_many){
    int i, j;
    int go_on;

    for (i = 0; i < how_many; i++)
    {
        print_array(how_many, data, "\nInside bubble\n");

        printf("i = %d, input any int to continue: \t", i);
        scanf("%d", &go_on);
        for (j = how_many - 1; j > i; j--)
        {
            if (data[j - 1] > data [j])
            {
                swap(&data[j - 1], &data[j]);
            }
        }
    }
}

int main (void){
    const int SIZE = 5;
    int grades [] = {78, 67, 92, 83, 88};

    print_array(SIZE, grades, "My grades: \n");
    printf("\n\n");
    bubble(grades, SIZE);
    printf("\n\n");
    print_array(SIZE, grades, "My sorted grades: \n");
    printf("\n\n");

    return 0;
}