/*Cuenta el numero de caracteres de un texto a traves de if y while
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    int blanks = 0, digits = 0, letters = 0, others = 0;
    int c;
    while ((c = getchar()) != EOF)
    {
        if (c == ' ')
        {
            ++ blanks;
        }
        else if (c >= '0' && c <= '9')
        {
            ++ digits;
        }
        else if (c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z')
        {
            ++ letters;
        }
        else
            ++ others;
    };
    printf("\nBlanks = %d, digits = %d, letters = %d, others = %d\n", blanks, digits, letters, others);
    return 0;
    
}