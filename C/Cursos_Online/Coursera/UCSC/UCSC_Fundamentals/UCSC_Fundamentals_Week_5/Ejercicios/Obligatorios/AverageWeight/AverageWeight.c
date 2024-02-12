/*Codigo que obtiene el promedio de datos obtenidos en un archivo de pesos de leones marinos, generando un arreglo e imprimiendo el promedio
 * Juan Luis Rivera
 * 17 de Enero 2024*/

/*Code in which we obtain data from an external file, on this code we sacrifice time in exchange 
* of space saved by creating a dinamicall array based on the size of the elements on the file instead of creating a fixed array
* large enough to gather all data but at the expense of reading two times the data file.*/

#include <stdio.h>

int get_SIZE(const char* data){ //Function that retrieves data from file
    int SIZE = 0, i = 0; //Initialize variable that will count the number of elements

    FILE *filename = fopen(data, "r"); //We open the file

    if (filename == NULL) { //Tries to find out if the file opened successfully

        fprintf(stderr, "Error openening the file!\n"); //Prints an error message

        return -1; //Returns error
    }

    else{ //Obtains the number of elements in the array

        while (fscanf(filename, "%d\t", &i) != EOF){ //Reads everything until EOF
            SIZE++; 
        }

        fclose(filename); //We close the data file

        return SIZE; //Returns the number of elements
    }
}

void create_array (const char* data, int SIZE, int array[]){ //We read the file again in order to create the array
    int weights[SIZE]; //We create a new array to store the data

    FILE *filename = fopen(data, "r"); //We open the data file again

    for (int i = 0; i < SIZE; i++)
    {
        fscanf(filename, "%d\t", &weights[i]); //We assign each element on the array the respective value in the data file
    }

    fclose(filename); //We close the data file

    for (int i = 0; i < SIZE; i++)
    {
        array[i] = weights[i]; //Updates the original array with the new array values
    }
}

void print_average(int array[], int SIZE){ //Prints the average in the array given its size
    double sum = 0.0; //We definte and initialize the variable where we are going to save the total of the array addition

    for (int i = 0; i < SIZE; i++) //Obtains the total sum of the values in the array
    {
        sum = sum + array[i]; 
    }

    printf("Average: %.2f\n\n", sum / SIZE); //Returns the average
}

void print_array(int how_many, int data[], char *str){ //Prints the given array
    int i;
    printf ("%s", str);

    for (i = 0; i < how_many; i++)
    {
        printf ("%d\t", data[i]);
    }

    printf("\n\n");
}

int main(void)
{
    const char* filename = "data.txt"; 

    int SIZE = get_SIZE(filename);

    int arreglo[SIZE];

    create_array(filename, SIZE, arreglo);

    print_array(SIZE, arreglo, "Data: \n");

    print_average(arreglo, SIZE); 

    return 0;
}