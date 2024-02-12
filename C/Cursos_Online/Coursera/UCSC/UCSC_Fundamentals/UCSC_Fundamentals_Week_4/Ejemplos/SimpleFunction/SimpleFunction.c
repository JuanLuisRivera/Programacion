/*Codigo que muestra el funcionamiento de una funcion simple
 * Juan Luis Rivera
 * 14 de Enero 2024*/

#include <stdio.h>

void simplefunct(int count){ //Programa que imprime los numeros de manera regresiva hasta 0
    while (count >= 0)
    {
        printf("----\n");
        printf("%d", count);
        printf("\n----\n");
        count --;
    }
}

int main (void){

    int repeat;
    printf("Print how many numbers?\n");
    scanf("%d", &repeat); 
    simplefunct(repeat); //Llama a la funcion simplefunct

    return 0;
}