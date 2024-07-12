import Data.List -- Se importa el modulo, el cual trabaja exclusivamente con listas

a = intersperse '*' "Hola" -- La funcion "intersperse" toma un elemento y una lista y coloca el elemento entre todos los elementos de la lista, de manera que en este caso quedaria "H*o*l*a"
b = intersperse 1 [2, 4, 6] -- En este caso la lista resultante seria "[2, 1, 4, 1, 6]"

