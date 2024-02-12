/*Distancia de un maraton en kilometros
 * Juan Luis Rivera
 * 8 de Enero de 2024 */

#include <stdio.h>

int main(void){
    int miles = 26, yards = 365;
    double kilometers;

    kilometers = 1.609 * (miles + yards / 1760.0);
    printf("\nUn maraton es igual a %lf kilometros.\n\n", kilometers);
    return 0;
}
