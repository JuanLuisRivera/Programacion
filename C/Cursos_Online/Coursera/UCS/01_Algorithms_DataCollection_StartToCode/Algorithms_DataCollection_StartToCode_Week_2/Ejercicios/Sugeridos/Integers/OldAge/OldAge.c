/*Muestra la operacion basica de resta con los enteros
 * Juan Luis Rivera
 * 19 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>

#define YEAR_BORN 1962
#define YEAR_TODAY 2024

int main (int argc, char** argv){
    //We define the variable needed for the exercise and difference between "YEAR TODAY" and "YEAR BORN"
    int oldAge = YEAR_TODAY - YEAR_BORN; 

    //We print the result using the propet format specifier for integers "%d"
    printf("Old age: %d\n", oldAge);

    return (EXIT_SUCCESS);
}