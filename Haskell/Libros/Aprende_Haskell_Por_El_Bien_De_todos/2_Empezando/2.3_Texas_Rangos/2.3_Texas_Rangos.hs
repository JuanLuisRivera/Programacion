a = [1..10] --Se obtiene una lista de enteros desde el 1 hasta el 10 utilizando la notacion ".." que implica
--El rango en el cual estaran los datos

b = [2, 2..10] --Se obtiene una lista de enteros desde el 2 hasta el 10 especificando que la lista tendra un un incremento de 2 en 2

c = [10, 9..1] --Se obtiene una lista de enteros desde el 10 hasta el 1 especificando que la lista tendra un decremento de 1 en 1

d = [0.1, 0.3 .. 1] --El uso de rangos con numeros flotantes dara comportamientos erroneos, por lo que
 --No se recomienda utilizarlos juntos
e = [1..] --Se obtiene una lista de enteros desde el 1 hasta el infinito
f = take 13 [13, 26..]--Se obtiene una lista de enteros desde el 13 hasta el infinito, pero unicamente evalua los 13 primeros
 --Valores que le requerimos
g = take 10 (cycle [1, 2, 3]) --La funcion "cycle" toma una lista de elementos y la repite de manera infinita
 --Por lo que el resultado de tomar 10 elementos de la lista sera "[1, 2, 3, 1, 2, 3, 1, 2, 3, 1]"
h = take 10 (repeat 5) --La funcion "repeat" toma un elemento y genera una lista infinita de ese elemento, por lo que el 
 --Resultado de tomar 10 elementos de la lista sera "[5, 5, 5, 5, 5, 5, 5, 5, 5, 5]"