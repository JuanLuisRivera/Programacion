/*Ejemplo de como definir tipos de datos propios
 * Juan Luis Rivera
 * 21 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>


int main (int argc, char** argv){
    //We define the way we intend to use our datatype
    typedef struct Student Student;

    //We define a "Student" datatype
    struct Student
    {
        int number;
        float percent;
        char grade;
    };
    

    //We declare a "Student" datatype with a name "student_0" and assign the variables "number" = 10101010, "percent" = 89.5, "grade" = 'B'
    Student student_0 = {10101010, 89.5, 'B'}; //We assign the values to the variable "student_0" by using "{}" and filling it accordingly to the datatype variables

    //We print the information of the variable "student_0"
    printf("Student 0\n");
    printf("---------\n");
    printf("Number: %d\n", student_0.number);
    printf("Number: %f\n", student_0.percent);
    printf("Number: %c\n", student_0.grade);  //We use the apropiate format specifier "%c" to print the char value
    printf("---------\n");
    printf("\n\n\n");

    //We declare a "Student" datatype with a name "student_1"
    Student student_1;

    student_1.number = 222222; //We assign the value of "222222" variable "number" in "student_1" with the "." operator
    student_1.percent = 56.5; //We assign the value of "56.5" to the variable "percent" in "student_1" with the "." operator
    student_1.grade = 'F'; //We assign the value of "'F'" to th variable "grade" in "student_1" with the "." operator

    //We print the information of the variable "student_0"
    printf("Student 1\n");
    printf("---------\n");
    printf("Number: %d\n", student_1.number); //We use the apropiate format specifier "%d" to print the int value
    printf("Number: %f\n", student_1.percent); //We use the apropiate format specifier "%f" to print the float value
    printf("Number: %c\n", student_1.grade);  //We use the apropiate format specifier "%c" to print the char value
    printf("---------\n");


    return (EXIT_SUCCESS);
}