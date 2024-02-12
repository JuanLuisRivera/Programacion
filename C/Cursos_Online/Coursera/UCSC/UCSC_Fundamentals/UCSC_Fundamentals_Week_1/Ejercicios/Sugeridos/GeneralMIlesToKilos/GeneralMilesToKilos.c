/*Distancia de un maraton en kilometros
 * Juan Luis Rivera
 * 10 de Enero de 2024 */

#include <stdio.h>

int main(void){
    double miles, yards, kilometers;

    printf("Enter miles + yards in that order:\n");
    scanf("%lf%lf", &miles, &yards);

    kilometers = 1.609 * (miles + yards / 1760.0);
    printf("\nLa distancia %lf en millas con %lf yardas es equivalente a %lf kilometros.\n\n", miles, yards, kilometers);
    return 0;
}
