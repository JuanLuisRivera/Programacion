/*
* File: Mean_StandardDeviation_Arrays.c
* Author: Juan Luis Rivera
* Date: 28 de enero de 2024
* Description: Program that shows the way of using and operating with arrays
*/

#define _CRT_SECURE_NO_WARNINGS

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

int main(int argc, char** argv)
{
    int scores[4];
    int size = sizeof(scores) / sizeof(scores[0]); //We calculate the size of the array by uysing the sizeof function

    for (int i = 0; i < 4; i++) //We insert data into the array
    {
        printf("Enter score %d: ", i + 1);
        scanf("%d", &scores[i]);
        while (scores[i] < 0 || scores[i] > 100)
        {
            printf("\n");
            printf("Invalid score. Enter score %d: ", i + 1);
            scanf("%d", &scores[i]);
        }
    }

    //We calculate the mean
    int sum = 0;
    for (int i = 0; i < size; i++)
    {
        sum += scores[i];
    }
    float mean = (float)sum / size;
    printf("Mean: %.2f\n", mean);

    //We calculate the standard deviation
    float sum_squared_differences = 0;
    for (int i = 0; i < size; i++)
    {
        sum_squared_differences += powf(scores[i] - mean, 2);
    }
    float standard_deviation = (float)sqrtf(sum_squared_differences / (size - 1));
    printf("Standard deviation: %.2f\n", standard_deviation);

    return (EXIT_SUCCESS);
}
