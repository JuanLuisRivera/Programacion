/*Codigo que regresa la tabla de senos y cosenos
 * Juan Luis Rivera
 * 13 de Enero 2024*/

#include<stdio.h>
#include<math.h> /* has  sin(), abs(), and fabs() */
int main(void)
{ 
    double interval;
    int i;
    for(i = 1; i <100; i++) //Its so the interval values are open starting from < 0 and until > 1
    {
        interval = i/30.0;
        printf("sin( %lf ) = %lf \t", interval, sin(interval));
        printf("cos( %lf ) = %lf \n", interval, cos(interval));
    }

    return 0;
}