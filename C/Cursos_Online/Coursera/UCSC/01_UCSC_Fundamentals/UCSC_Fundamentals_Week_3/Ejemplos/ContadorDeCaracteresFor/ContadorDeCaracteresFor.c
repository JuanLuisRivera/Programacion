/*Cuenta el numero de caracteres de un texto a traves de if y for
 * Juan Luis Rivera
 * 10 de Enero 2024*/

#include <stdio.h>

int main(void){
    int blanks = 0, digits = 0, letters = 0, others = 0, total_chars = 0;
    int c;
    for (; (c = getchar()) != EOF; total_chars++)
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
    }
    printf("\nBlanks = %d, digits = %d, letters = %d, others = %d, total characters = %d.\n", blanks, digits, letters, others, total_chars);
    return 0;
    
}