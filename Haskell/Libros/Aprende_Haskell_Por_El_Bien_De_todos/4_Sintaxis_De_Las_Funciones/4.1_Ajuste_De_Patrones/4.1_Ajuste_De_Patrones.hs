--El ajuste de patrones consiste en especificar pautas que deben ser seguidas por los datos.
lucky :: (Integral a) => a -> String --Definimos una funcion que pertenece a la clase "Integral" revisada anteriormente,
--La cual toma un numero y regresa una cadena
lucky 7 = "Suerte" --Especificamos que si el numero es el numero "7" se imprimira una cadena especifica
lucky x = "No es de la suerte" --Si el numero es distinto, se retorna una cadena diferente
--Es importante que los patrones se verifican de arriba hacia abajo

sayMe :: (Integral a) => a -> String --Se define una funcion que evalua diferentes numeros y en caso de ajustarse a un patron especifico imprimen una
--Cadena acorde al patron definido
sayMe 1 = "¡Uno!"
sayMe 2 = "¡Dos!"
sayMe x = "No es 1 ni 2" --El caso general es importante porque en caso de no implementarlo se producira un error en caso de que algun valor
--No coincida con los patrones definidos

--De esta manera se crean dos funciones en las cuales se ocupa el ajuste de patrones

factorial :: (Integral a) => a -> a
factorial 0 = 1
factorial x = factorial(x - 1) * x 
--Se define la funcion factorial de manera recursiva

addVectors :: (Num a) => (a, a) -> (a, a) -> (a, a)
addVectors (a, b) (c, d) = (a + c, b + d)
--Se define una funcion que suma dos vectores en 2D en forma de tupla

first :: (a, b, c) -> a
first (a, _, _) = a --El operador "_" se ocupa cuando no importa el tipo de dato que se ocupa en ese lugar

second :: (a, b, b) -> b
second (_, b, _) = b

third :: (a, b, c) -> c
third (_, _, c) = c
--Se definen tres funciones para trabajar con triplas de manera similar a como funcionan las funcion "fst" y "snd" con duplas

tell :: (Show a) => [a] -> String
tell [] = "La lista esta vacia"
tell (x:[]) = "Un elemento: " ++ show x
tell (y:x:[]) = "Dos elementos: " ++ show x ++ show y
tell (y:x:_) = "Mas de dos elementos: " ++ show x ++ show y
--Se define una funcion la cual muestra los primeros elementos en una lista

lenght' :: (Num b) => [a] -> b
lenght' [] = 0
lenght' (_: xs) = 1 + lenght' xs
--Se define la funcion que nos dice la longitud de una lista de manera recursiva

capital :: String -> String
capital "" = "Vacia"
capital all@(x:_) = "La primera letra de " ++ all ++ " es " ++ [x] --Se utiliza el operador "@" para acceder a la lista completa
--Se define una funcion que regresa una lista y el primer elemento de la lista