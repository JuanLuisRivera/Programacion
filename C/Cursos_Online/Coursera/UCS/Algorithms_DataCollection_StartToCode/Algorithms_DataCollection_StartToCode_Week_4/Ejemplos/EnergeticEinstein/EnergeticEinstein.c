/*Programa que calcula la distancia entre 2 puntos en dos dimensiones
 * Juan Luis Rivera
 * 22 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

//We define a macro for the constant speed of light 
#define SPEED_LIGHT 300000

int main (int argc, char** argv){
    //We declare a variable of datatype "float" with id "mass" that will store the data the user inputs
    float mass;

    //We declare a variable of datatype "float" with id "energy" that will store the computation we do with the variable mass
    float energy; 


    //Prompts for requesting the user data
    printf("Enter mass of object in kg: \n");
    scanf("%f", &mass);

    //Print the energy given the data stored in "mass"
    energy = mass * powf(SPEED_LIGHT, 2);

    printf("Energy in the object = %f\n", energy);

    return (EXIT_SUCCESS);
}