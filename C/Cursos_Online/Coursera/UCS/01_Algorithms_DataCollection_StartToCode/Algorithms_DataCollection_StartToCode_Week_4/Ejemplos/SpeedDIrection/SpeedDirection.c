/*Programa que calcula la velocidad, la direccion entre dos puntos y un vector de velocidad
 * Juan Luis Rivera
 * 22 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

#define TIME_TO_MOVE 3

#ifndef M_PI
    #define M_PI 3.141592653589933846
#endif

int main (int argc, char** argv){
    //We define a variable of type float where to store the data from the points given by the user
    float pointOne_x;
    float pointOne_y; 

    float pointTwo_x;
    float pointTwo_y;

    typedef struct Vector Vector;

    struct Vector
    {
        float x;
        float y;
    };
    


    //Prompts for requesting the user data
    printf("Enter x for first point:");
    scanf("%f", &pointOne_x);

    printf("Enter y for first point:");
    scanf("%f", &pointOne_y);

    printf("Enter x for second point:");
    scanf("%f", &pointTwo_x);

    printf("Enter y for second point:");
    scanf("%f", &pointTwo_y);

    printf("\n");
    
    float delta_X = pointTwo_x - pointOne_x;
    float delta_Y = pointTwo_y - pointOne_y;

    Vector velocity;
    velocity.x = delta_X / TIME_TO_MOVE;
    velocity.y = delta_Y / TIME_TO_MOVE;

    printf("Velocity vector: (%.2f, %.2f)\n", velocity.x, velocity.y);
 
    float distancia = sqrtf(powf(delta_X, 2) + powf(delta_Y, 2));


    float speed = distancia / TIME_TO_MOVE;


    float direccion = atan2f(delta_Y, delta_X);
    direccion = direccion * 180 / M_PI;

    Vector unitDireccion;
    unitDireccion.x = delta_X;
    unitDireccion.y = delta_Y;

    float length = sqrtf(powf(unitDireccion.x, 2) + powf(unitDireccion.y, 2));
    unitDireccion.x = unitDireccion.x * speed;
    unitDireccion.y = unitDireccion.y * speed;
    printf("Vector de velocidad: (%.2f, %.2f)\n", velocity.x, velocity.y);


    printf("Velocidad: %.2f\n", speed);

    printf("Direccion: %.2f\n", direccion);

    return (EXIT_SUCCESS);
}