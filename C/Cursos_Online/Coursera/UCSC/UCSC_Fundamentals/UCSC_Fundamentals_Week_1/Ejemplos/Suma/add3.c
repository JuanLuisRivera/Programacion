/*Read tree float and print the result*/
/*8 de Enero del 2024*/
#include <stdio.h>

int main(void){
    float a, b, c, sum;
    printf("Input tree floats:\n");
    scanf("%f%f%f", &a, &b, &c);
    printf("a = %f, b = %f, c = %f\n", a, b, c);
    sum = a + b + c;
    printf("sum = %f\n\n", sum);
    return 0;
}
