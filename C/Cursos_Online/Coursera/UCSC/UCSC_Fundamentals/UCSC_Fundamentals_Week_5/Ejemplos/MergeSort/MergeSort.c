/*Codigo que muestra el funcionamiento del algoritmo merge y merge sort en C
 * Juan Luis Rivera
 * 15 de Enero 2024*/

#include <stdio.h>

void print_array(int how_many, int data[], char *str){
    int i;
    printf ("%s", str);

    for (i = 0; i < how_many; i++)
    {
        printf ("%d\t", data[i]);
    }
}

void merge (int a[], int b[], int c[], int how_many){
    int i = 0, j = 0, k = 0;
    while (i < how_many && j < how_many)
    {
        if (a[i] < b[j])
        {
            c[k++] = a[i++];
        }

        else{
            c[k++] = b[j++];
        }
    }
    while (i < how_many)
    {
        c[k++] = a[i++];
    }
    while (j < how_many)
    {
        c[k++] = b[j++];
    }
}

void merge_sort(int key[], int how_many){
    int j, k;
    int w[how_many];

    for (k = 1; k < how_many; k *= 2)
    {
        for (j = 0; j < how_many - k; j += 2 * k)
        {
            merge(key + j, key + j + k, w + j, k);
        }
        for (j = 0; j < how_many; j++)
        {
            key[j] = w[j];
        }
    }
}

int main (void){
    const int SIZE = 5;
    const int SIZE_D = 8;
    int a[] = {67, 82, 83, 88, 99};
    int b[] = {58, 69, 72, 81, 88};
    int d[] = {46, 34, 94, 424, 24, 67, 83, 22};
    int c[2 * SIZE];

    print_array(SIZE, a, "My grades: \n");
    printf("\n\n");

    print_array(SIZE, b, "More Grades: \n");
    printf("\n\n");

    merge(a, b, c, SIZE);

    print_array(2 * SIZE, c, "My sorted grades\n");
    printf("\n\n");

    merge_sort(d, SIZE_D);

    print_array(SIZE_D, d, "Merge sort use: \n");
    printf("\n\n");

    return 0;
}