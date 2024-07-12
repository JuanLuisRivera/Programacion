map' :: (a -> b) -> [a] -> [b] -- Se define una funcion que toma un argumento "a" y devuelve un argumento "b", donde posteriomente se aplicara
-- Sobre una lista "[a]" y regresara una lista "[b]"
map' _ [] = []
map' f (x:xs) = f x : map' f xs -- Se define la funcion "f", la cual se aplicara al elemento "x" perteneciente a la lista "xs" y se agregara
-- De manera recursiva a la lista de los elementos restantes de la lista "xs"

a = map (+ 2) [1, 2, 3, 4] -- Se aplicara la funcion "(+ 2)" a todos los elementos de la lista, retornando la lista [3, 4, 5, 6]

filter' :: (a -> Bool) -> [a] -> [a] -- Se define la funcion que toma un elemento de tipo "a" y regresa un "Bool", la cual se aplica sobre
-- Una lista de elementos de tipo "[a]" y nos regresa una lista de elementos de tipo "a"
filter' _ [] = []
filter' p (x: xs) -- Se define la funcion "p" la cual se aplicara a los elementos de la lista de tipo "a"
    | p x = x : filter' p xs -- En caso de que la aplicacion de "p" en "x" sea "True", se agregara a la lista
    | otherwise = filter p xs -- En cualquier otro caso, se procedera con el siguiente elemento de la lista "xs"

b = filter (> 3) [1, 2, 3, 4, 5, 6, 7, 8] -- Se aplicara la funcion (> 3) a todos los elementos de la lista, produciendo [4, 5, 6, 7, 8]

-- Nota: El uso de map y filter es un caso que mejora la legibilidad, sin embargo es perfectamente posible ocupar listas por comprension

quicksort :: (Ord a) => [a] -> [a] --Se reescribe la funcion "quicksort" pero ocupando filter para mostrar la mejora en la legibilidad
quicksort [] = []
quicksort (x:xs) =
    let smallerSorted = quicksort (filter (<=x) xs)
        biggerSorted = quicksort (filter (>x) xs)
    in  smallerSorted ++ [x] ++ biggerSorted

largestDivisible :: (Integral a) => a
largestDivisible = head (filter p [1000000, 999999 ..]) -- Se toma el primer elemento (el mas grande) de la lista de numeros que cumplen
-- La condicion dada por la funcion "p"
    where p x = x `mod` 3829 == 0 -- Se define la funcion "p" para el filtrado, donde se necesita cumplir que el elemento "x"
    -- Sea divisible por "3829"

c = sum (takeWhile (< 10000) (filter odd (map (^2)[1 ..]))) -- Se aplicara la funcion (^2) a traves de "map" a todos los elementos de la
-- Lista infinita empezando por "1" que cumplan la condicion de ser impares, generando a su vez otra lista la cual se detendra cuando la 
-- Los elementos de la nueva lista sean mayores a "10000"

d = sum (takeWhile (<10000) [n^2 | n <- [1..], odd (n^2)]) -- Se reescribe la funcion ocupando listas por comprension para comparar la
-- Legibilidad entre ambas formas de resolver el mismo problema

chain :: (Integral a) => a -> [a] -- Se crea una funcion que nos da los elementos de la serie de Collatz
chain 1 = [1]
chain n
    | even n =  n:chain (n `div` 2)
    | odd n  =  n:chain (n*3 + 1)

numLongChains :: Int
numLongChains = length (filter isLong (map chain [1..50])) -- Establecemos que se aplicara la funcion "chain" a la lista de "1" a "100"
-- Las cuales cumplen la propiedad dada por la funcion "isLong" y posteriormente se hace el conteo de todos los elementos en la lista
-- Que obtenemos al ocupar la funcion "filter" a traves de la funcion "length"
    where isLong xs = length xs > 15 -- Definimos la funcion "isLong" donde cualquier secuencia mayor a 15 sera verdadera