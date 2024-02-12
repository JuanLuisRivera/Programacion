/*Muestra el funcionamiento basico del tipo de dato flotante
 * Juan Luis Rivera
 * 21 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>

//We define a macro
//We define a macro to add readibility to the code
#define MAX_SCORE 100000

int main (int argc, char** argv){
    //We define a variable type "int" with a name "score" and initializate it with a value of "4586"
    int score = 4586;

    //We define a variable type "float" with a name "percent" and initializate it with the result of "score / MAX_SCORE"
    float percentage = (float) score / MAX_SCORE; //We force the "float" division by doing a typecast on one of the variables

    //We print the score using the proper format specifier for integers "%d"
    printf("Score: %d\n", score);

    //We print the percentaje using the proper format specifier for float "%f" and 0 decimal places ".0"
    printf("Score: %.0f\n", percentage);

    //We print the percentaje using the proper format specifier for float "%f" and two decimal places ".2"
    printf("Score: %.2f\n", percentage);

    return (EXIT_SUCCESS);
}