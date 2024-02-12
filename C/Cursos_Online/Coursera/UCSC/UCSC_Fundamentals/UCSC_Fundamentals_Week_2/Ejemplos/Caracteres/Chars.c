/*Muestra el funcionamiento de characteres y ASCII
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    char c = 'a';
    printf("c en ASCII es: %d\n", c);
    printf("c en caracters es: %c\n", c);
    printf("Tres characteres consecutivos son: %c%c%c \n", c, c + 1, c + 2);
    printf("Tres campanas consecutivas son: %c%c%c \n", '\a', '\a', '\a');
    return 0;
}