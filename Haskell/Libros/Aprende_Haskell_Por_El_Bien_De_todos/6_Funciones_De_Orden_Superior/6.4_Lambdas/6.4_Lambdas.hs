chain :: (Integral a) => a -> [a] -- Se crea una funcion que nos da los elementos de la serie de Collatz
chain 1 = [1]
chain n
    | even n =  n:chain (n `div` 2)
    | odd n  =  n:chain (n*3 + 1)

numLongChains :: Int
numLongChains = length (filter (\xs -> length xs > 15) (map chain [1..100])) --Hacemos uso de las lambdas en lugar de la expresion "where"
-- Definimos una funcion lambda a traves del operador "\" y la rodeamos entre parentesis "()" para evitar que se extiendan
-- Las lambdas al igual que "if" son expresiones, de manera que se pueden ocupar en cualquier lugar

-- Es importante notar que el unico proposito de las lambdas es actuar como funciones anonimas que solo se ocuparan una sola vez
-- En funciones de orden superior, de manera que es erroneo ocupar expresiones lambda en cualquier expresion si dificultan la 
-- Legibilidad del codigo

a = zipWith (\a b -> (a * 30 + 3) / b) [5,4,3,2,1] [1,2,3,4,5] -- Se define una lambda que toma dos parametros, la cual se pasara
-- Como argumento a la funcion "zipWith", lo que dara como resultado la lista [153, 61.5, 31, ...]

flip' :: (a -> b -> c) -> b -> a -> c -- Se define la funcion "flip'" a traves de la funcion lambda 
flip' f = \x y -> f y x