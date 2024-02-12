/*
 * Circulo y area
 * 8 de Enero del 2024
 * */

#include <stdio.h>

#define PI 3.14159

int main(){
    double area = 0.0, radius = 0.0;
    printf("Enter radius:\n");
    scanf("%lf", &radius);
    area = PI * radius * radius;
    printf("Radius of %lf meters \n Area is %lf sq. meters\n", radius, area);
    return 0;
}
