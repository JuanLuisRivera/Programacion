/*Calcula la diferencia de edades
 * Juan Luis Rivera
 * 19 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>

#define OLD_YEAR_BORN 1962
#define YEAR_TODAY 2024
#define MY_YEAR_BORN 1998

int main (int argc, char** argv){
    //We define the variable needed for the exercise and difference between "YEAR TODAY" and "YEAR BORN"
    int oldAge = YEAR_TODAY - OLD_YEAR_BORN; 

    //We define the value of my age
    int age = YEAR_TODAY - MY_YEAR_BORN;

    //We define the difference in a new variable
    int difference = oldAge - age;

    //We print the result using the propet format specifier for integers "%d"
    printf("Compared age: %d\n", difference);

    return (EXIT_SUCCESS);
}