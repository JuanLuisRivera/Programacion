/*Code that returns the value of the sine function between 0 and 1
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>
#include <math.h>

int main(void)
{
    double x = 0.0, result = 0.0;                           //Initialize and assigns value to the variables to use

    printf("Input the value of x between 0 and 1: \n");     //Request the value form the user
    scanf("%lf", &x);
    result = sin(x);                                        //Computes the value of the sin(x)
    printf("sin(%lf) = %lf\n", x, result);                  //Prints the value of x and the value of the function
    return 0;
}