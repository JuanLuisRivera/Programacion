/*Codigo que muestra el funcionamiento de las reglas para variables en C
 * Juan Luis Rivera
 * 14 de Enero 2024*/

#include <stdio.h>

extern int reps = 0;

void f(void){
    static int called = 0;
    printf("f called %d\n", called);
    called ++;
    reps = reps + called;
}

int main (void){
    auto int i = 1;
    const int limit = 10;
    for (i = 1; i < limit;  i++)
    {
        printf ("i local = %d, reps global = %d\n", i, reps);
        f();
    }

    return 0;
}