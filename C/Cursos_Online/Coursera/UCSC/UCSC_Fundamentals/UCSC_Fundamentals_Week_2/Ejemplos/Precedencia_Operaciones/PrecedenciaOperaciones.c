/*Muestra la forma en que se ejecutan las operaciones
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    int a = 5, b = 7, c = 0, d = 0;
    c = a - b;
    printf("a = %d, b = %d, c = %d, d = %d\n", a, b, c, d);
    c = a / b;
    d = b / a;
    printf("a = %d, b = %d, c = %d, d = %d\n", a, b, c, d);
    c = a % b;
    d = b % a;
    printf("a = %d, b = %d, c = %d, d = %d\n", a, b, c, d);
    c = -a - b;
    d = -b - a;
    printf("a = %d, b = %d, c = %d, d = %d\n", a, b, c, d);
    c = ++a + b++;
    d += 5;
    printf("a = %d, b = %d, c = %d, d = %d\n", a, b, c, d);
    return 0;
}