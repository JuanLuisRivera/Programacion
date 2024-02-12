/*Codigo que regresa la tabla de senos y cosenos
 * Juan Luis Rivera
 * 13 de Enero 2024*/

#include<stdio.h>
#include<math.h> /* has  sin(), abs(), and fabs() */
int main(void)
{ 
    double interval;
    int i;
    for(i = 0; i <30; i++) //Its proposed a 30 value interval
    {
        interval = i/10.0;
        printf("sin( %lf ) = %lf \t", interval, fabs(sin(interval)));
    }


    printf("\n+++++++\n");
    return 0;
}