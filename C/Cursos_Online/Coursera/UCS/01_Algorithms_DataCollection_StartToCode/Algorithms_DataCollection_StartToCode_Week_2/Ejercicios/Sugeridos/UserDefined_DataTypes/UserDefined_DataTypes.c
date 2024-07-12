/*Ejemplo de como definir tipos de datos propios
 * Juan Luis Rivera
 * 21 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>


int main (int argc, char** argv){
    //We define the way we intend to use our datatype
    typedef struct GPS GPS;

    //We define a "GPS" datatype
    struct GPS
    {
        float longitude; //We declare one of its "members" or "fields" is a float an has a name "longitude"
        float latitude; //We declare one of its "members" or "fields" is a float an has a name "latitude"
    };
    

    //We declare a "GPS" datatype with a name "actual" and assign the fields "longitude" = 0.55, "latitude" = 17.3, we also use the "f" right next to a number to specify that its a float and not a double
    GPS actual = {0.55f, 17.3f}; //We assign the values to the variable "actual" by using "{}" and filling it accordingly to the datatype variables

    //We declare a "GPS" datatype with a name "visit" and assign the fields "longitude" = 33.19, "latitude" = 47.59, we also use the "f" right next to a number to specify that its a float and not a double
    GPS visit = {33.19f, 47.59f}; //We assign the values to the variable "visit" by using "{}" and filling it accordingly to the datatype variables

    //We print the information of the variable "student_0"
    printf("Difference between distance: \n");
    printf("---------\n");
    printf("Difference in longitude: %f\n", visit.longitude - actual.longitude);
    printf("Difference in latitude: %f\n", visit.latitude - actual.latitude);
    printf("\n\n");


    return (EXIT_SUCCESS);
}