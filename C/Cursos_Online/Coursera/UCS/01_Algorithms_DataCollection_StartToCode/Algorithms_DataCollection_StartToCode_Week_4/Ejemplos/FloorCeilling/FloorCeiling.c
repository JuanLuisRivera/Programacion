/*Se programa la funcion escalon y piso para numeros positivos
 * Juan Luis Rivera
 * 22 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>


int main (int argc, char** argv){
    //We define a variable of type float where to store the data from the user
    float x; 

    printf("Enter floating point number: \n");
    scanf("%f", &x);

    //We make use of a typecast in order to round the value down
    printf("Floor: %d\n", (int)x);

    //We make use of the typecast again but this time we have to make sure the new value is either equal or bigger than the next integer
    printf("Ceilling: %d\n", (int)(x + 0.99)); //As we are using 0.99 there is a limitation on the number of decimals we can use

    //We print the result using the mathlib.h "floorf" function
    printf("Floor using math.h: %d\n", (int)floorf(x));

    //We print the result using the mathlib.h "ceilf" function
    printf("Ceiling using math.h: %d\n", (int)ceilf(x));

    return (EXIT_SUCCESS);
}