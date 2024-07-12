/*
* File: Min_Max_User-Fixed_Size.c
* Author: Juan Luis Rivera
* Date: 28 de enero de 2024
* Description: Exercise 1 and 2: where we allocate data input form the user by creating a fixed size array and fined the maximum and minimum 
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv)
{
    int array [5];  //Exercise 1, we create a fixed size array of 5 elements 

    int max, min, aux; //We create variables to analyze the maximum and minimum
    
    for (int i = 0; i < 5; i++) //For loop to insert data into the array
    {
        printf("Enter score (from 0 - 100) %d: ", i + 1); //We indicate the number of element that will be inserted
        scanf("%d", &array[i]);

        while (array[i] < 0 || array[i] > 100) //We validate that the user data input
        {
            printf("\n");
            printf("Invalid score. Enter score %d: ", i + 1);
            scanf("%d", &array[i]);
        }
    }

    for (int i = 0; i < 5; i++) //For loop to obtain the maximum and minimum in the array
    {
        aux = array[i]; //We use the auxiliary variable to store the value we are going to use to compare the maximum and minimum

        if (i == 0) //We assign values to the variables in the first iteration
        {
            max = array[i]; //We assign the first element in the array to be the maximum
            min = array[i]; //We assign the first element in the array to be the minimum
        }
        
        if (aux > max)
        {
            max = aux;
        }

        else if (aux < min)
        {
            min = aux;
        }
    }

    printf("Maximum: %d\n", max); //We print the maximum
    printf("Minimum: %d\n", min); //We print the minimum    
    
    return (EXIT_SUCCESS);
}