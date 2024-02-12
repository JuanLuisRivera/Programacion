/*Ejemplo de impresion de numeros flotantes y sus especificadores
 * Juan Luis Rivera
 * 20 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>


int main (int argc, char** argv){
    //We define an int variable
    int hoursDriven = 25; 

    //We define a second int variable
    int milesDriven = 262;

    //We define a new variable called "miles per hour" that will have decimal representation
    //We also make a typecast on "milesDriven" or "hoursDriven" in order to force the system to make a floating division instead of the integer division
    float mph = (float) milesDriven / hoursDriven;

    //We print the variable "mph" with its respectiva specifier, in the case of floats, %f
    printf("MPH: %f\n", mph);

    //We print the variable "mph" again but with the desired lenght of decimal places by specifying with "%.xf" where x is the number of decimal places we desire
    printf("MPH: %.5f\n", mph);

    //We print again but now declaring the spacing size
    printf("MPH in 00 spaces: %0.5f\n", mph);
    printf("MPH in 15 spaces: %15.5f\n", mph);
    printf("MPH in 25 spaces: %25.5f\n", mph);
    printf("MPH in 35 spaces: %35.5f\n", mph);

    return (EXIT_SUCCESS);
}