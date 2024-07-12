sum' :: (Num a) => [a] -> a
sum' xs = foldl (\acc x -> acc + x) 0 xs -- Se define la funcion "sum'", la cual suma todos los elementos de una lista
-- Para ello tomamos la lista "xs" y aplicamos la funcion "foldl", la cual se aplica a una expresion lambda "\acc x -> acc + x"
-- Esta expresion lambda toma dos argumentos "acc" y "x" y da como resultado "acc + x", a esta funcion le pasamos como argumentos
-- Iniciales el valor "0" correspondiente a la variable "acc" y el elemento de la lista "xs" correspondiente a la variable "x"

a = sum' [1, 2, 3, 4, 5]
-- En primera instancia la variable comienza con la variable "acc" igual a 0, y la variable "x" igual a "1" de manera que el valor 
-- Obtenido por la funcion "foldl" en esta primera instancia es "0 + 1 = 1", en la segunda instancia el valor de "acc" sera igual al 
-- Obtenido por "foldl" en la primera instancia, de manera que "acc" = 1 y "x" = 2, de manera que "foldl" = 3
-- En la tercera instancia el valor de "acc" sera el del ultimo resultado de "foldl", es decir "acc" = 3 y "x" = 3, por tanto "foldl" = 6
-- Y asi de manera sucesiva hasta obtener el resultado de 15

sum'' :: (Num a) => [a] -> a
sum'' = foldl (+) 0 -- Se reescribe la funcion de forma currificada, omitiendo la expresion lambda al ser equivalente a la funcion "+"

elem' :: (Eq a) => a -> [a] -> Bool
elem' y ys = foldl (\acc x -> if x == y then True else acc) False ys -- Se define la expresion lambda que toma dos argumentos "acc" y "x"
-- Si determina que el elemento "y" es igual a un elemento de la lista, regresa un valor "True", en caso contrario regresa el valor de "acc" 
-- Y avanza el elemento de la lista un elemento, el cual se asignara a la variable "x" y se volvera a comparar con
-- El valor de la variable "y", esta manera se comparara la variable "y" con todos los elementos de la lista 

sum''' :: (Num a) => [a] -> a
sum''' xs = foldr (\x acc -> acc + x) 0 xs -- Se reescribe la funcion "sum'" ocupando pliegues por la derecha

b = sum''' [1, 2, 3, 4, 5] -- Se empieza la suma con la operacion "5 + acc", donde "acc" tiene un valor de 0, el resultado se almacena en la variable "acc",
-- De manera que en la siguiente iteracion, el valor de "acc" sera 5 y la operacion sera "4 + 5", almacenando el valor de nuevo, en la variable "acc",
-- De forma que el resultado final sera "15", el cual seria el mismo que si hubieramos utilizado pliegues por la izquierda.

map' :: (a -> b) -> [a] -> [b]
map' f xs = foldr (\x acc -> f x : acc) [] xs -- Se reecribe la funcion "map" ocupando pliegues por la derecha "foldr", en este caso
-- El argumento que conservara la informacion del elemento de la lista sera el argumento derecho, por lo cual se cambia el orden con
-- Respecto las funciones anteriores que ocupaban pliegues por la izquierda "foldl", de manera que en la primera iteracion
-- "x" tendra el valor del elemento de la lista desde el lado derecho, y "acc" sera "[]", de manera que se aplicara primero la funcion
-- "f" sobre el elememento derecho de "xs" y posteriormente se agregara a la lista de la variable "acc"

c = map' (+3) [1, 2, 3, 4, 5] -- Al aplicar la funcion "map'" el orden en que se recorre la lista es empezando por el lado derecho,
-- Por lo tanto se tomara el ultimo elemento de la lista y se asignara a la variable "x", de manera que tendra un valor inicial de "5"
-- Posteriormente se le aplicara la funcion "f" (+3), dando como resultado "8" y posteriormente se concatenara a la lista de la variable
-- "acc", la cual se inicializo como "[]", quedando entonces la lista final como [4, 5, 6, 7, 8]

-- Es importante notar que cuando se ocupan pliegues por la izquierda sobre listas infinitas nunca se alcanzara el final, mientras que 
-- Si se ocupan pliegues por la derecha sobre listas infinitas si tendran final, ya que acotamos la lista al realizar el pliegue

maximum' :: (Ord a) => [a] -> a
maximum' = foldr1 (\x acc -> if x > acc then x else acc) --Ejemplo de la implementacion de la funcion maximum ocupando pliegues derechos
-- La funcion "foldr1" es similar a la funcion "foldr" solo que no necesitan indicarse valores de inicio, ya que asumen que el primer
-- Elemento de la lista es el valor de inicio y posteriomente realizan el pliegue

reverse' :: [a] -> [a]
reverse' = foldl (\acc x -> x : acc) [] --Implementacion de la funcion reverse' ocupando pliegues izquierdos

product' :: (Num a) => [a] -> a
product' = foldr1 (*) 

filter' :: (a -> Bool) -> [a] -> [a]
filter' p = foldr (\x acc -> if p x then x : acc else acc) [] -- 

head' :: [a] -> a
head' = foldr1 (\x _ -> x)

last' :: [a] -> a
last' = foldl1 (\_ x -> x) --

d = scanl (+) 0 [1, 2, 3, 4, 5, 6] --

e = scanr (+) 0 [1, 2, 3, 4, 5, 6] --