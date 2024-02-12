/*Muestra el funcionamiento basico de operaciones de tipo float
 * Juan Luis Rivera
 * 21 de Enero 2024*/

#include <stdio.h>
#include <stdlib.h>

//Line that allows to use the "scanf" function
#define _CRT_SECURE_NO_WARNINGS 

int main (int argc, char** argv){
    //We define three variables of type "float" where we will store diferent 
    float originalFahren = 0; //Variable to store the original fahrenheit value initialized with a value of "0"
    float converteCelsiu = 0; //Variable to store the celsius converted value from the original fahrenheit value initialized with a value of "0"
    float backToFahrenhe = 0; //Variable to store the converted fahrenheit value from the celsius converted value initialized with a value of "0"
    
    //We print a message to request the user the value of the temperature in fahrenheit
    printf("Enter temperature (Fahrentheit):");

    //We read the value input from the keyboard and assign it to the "originalFahren" variable
    scanf("%f", &originalFahren); //The symbol "&" refers to the address of the variables in memory and because of the scanf implementation it is needed to assign the value to the variable


    //We assign the value of the convertion operation from fahrenheit to celsius to the "converteCelsiu" variable
    converteCelsiu = (originalFahren - 32) / 9 * 5;

    //We assign the value of the convertion operation from celsius to fahrenheit to the "backToFahrenhe" variable
    backToFahrenhe = converteCelsiu * 9 / 5 + 32;

    //We print the values of each variable using the proper format specifier for float "%f"
    printf("Original Fahrenheit temperature: %f\n", originalFahren);
    printf("Converted to Celsius temperature: %f\n", converteCelsiu);
    printf("Converted back to Fahrenheit temperature: %f\n", backToFahrenhe);

    return (EXIT_SUCCESS);
}