/*Muestra el funcionamiento de un flujo for simple
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    int i = 0;
    for (int i = 0; i < 5; i++)
    {
        printf("%d\n", i);
    }
    printf("\nFor inverso.\n\n");
    for (int i = 5; i > 0; i--)
    {
        printf("%d\n", i);
    }
    
    return 0;
    
}