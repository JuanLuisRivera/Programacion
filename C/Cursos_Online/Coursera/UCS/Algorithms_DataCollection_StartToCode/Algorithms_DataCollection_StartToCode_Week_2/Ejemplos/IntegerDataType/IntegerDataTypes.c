/*Muestra el funcionamiento basico del tipo de dato entero
 * Juan Luis Rivera
 * 19 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>

//We define a macro to add readibility to the code
#define MINUTES_PER_HOUR 60

int main (int argc, char** argv){
    int totalMinutes = 130; 

    //Calculate and print hours
    int hours = totalMinutes / MINUTES_PER_HOUR;

    //We print the result using the proper format specifier for integers "%d"
    printf("Hours: %d\n", hours);

    //Calculate and print minutes
    int minutes = totalMinutes % MINUTES_PER_HOUR;
    printf("Minutes: %d\n", minutes);

    return (EXIT_SUCCESS);
}