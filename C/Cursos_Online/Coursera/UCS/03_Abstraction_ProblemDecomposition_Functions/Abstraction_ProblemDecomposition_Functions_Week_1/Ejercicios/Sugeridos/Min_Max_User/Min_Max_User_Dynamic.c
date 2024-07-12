/*
* File: Min_Max_User_Dynamic.c
* Author: Juan Luis Rivera
* Date: 28 de enero de 2024
* Description: Exercise 3: Program in which we use dynamic memory allocation and we print the maximum and minimum value of the array
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv)
{

    int n = 0; //Variable to obtain the number of elements that will be stored in the array

    /*Further explanation: We initialize it to 0 in order to void the undefined behavior*/

    int max, min, aux; //Variables to analyze the maximum and minimum

    printf("Enter the number of elements to store: \n"); //We request the user the number of elements that will be stored in the array

    scanf("%d", &n); //We assign the number of elements that will be stored in the array to the variable "n"
    
    int* scores = malloc(n * sizeof(int)); //We create the pointer to the array and initilize it with the number of elements that will be stored in the array
    
    /*Further explanation: We use malloc() in order to reserve the ammount of memory we will be using, in this case n * the size of an integer*/

    for (int i = 0; i < n; i++) //For loop to insert data into the array
    {
        printf("Enter score (from 0 - 100) %d: ", i + 1); //We indicate the number of element that will be inserted
        scanf("%d", &scores[i]);

        while (scores[i] < 0 || scores[i] > 100) //We validate that the user data input
        {
            printf("\n");
            printf("Invalid score. Enter score %d: ", i + 1);
            scanf("%d", &scores[i]);
        }
    }

    for (int i = 0; i < n; i++) //For loop to obtain the maximum and minimum in the array
    {
        aux = scores[i]; //We use the auxiliary variable to store the value we are going to use to compare the maximum and minimum

        if (i == 0) //We assign values to the variables in the first iteration
        {
            max = scores[i]; //We assign the first element in the array to be the maximum
            min = scores[i]; //We assign the first element in the array to be the minimum
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

    free(scores); //We free the memory use to store the array

    scores = NULL; //We set the pointer to NULL in order to avoid a stale pointer
    
    return (EXIT_SUCCESS);
}