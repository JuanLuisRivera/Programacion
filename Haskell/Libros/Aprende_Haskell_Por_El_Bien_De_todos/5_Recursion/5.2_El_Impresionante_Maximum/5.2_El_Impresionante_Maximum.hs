maximum1' :: (Ord a) => [a] -> a -- Se tomaran elementos que pueden ser ordenados
maximum1' [] = error "Máximo de una lista vacía" --Caso base 1
maximum1' [x] = x --Caso base 2
maximum1' (x:xs) --Algoritmo que imprime el eloemento mas grande en una lista
    | x > maxTail = x -- Se comparan todos los elementos de la lista y se regresa el mayor
    | otherwise   = maxTail
    where maxTail = maximum1' xs
-- Funcion que regresa el elemento mayor de una lista de elementos que se pueden ordenar

maximum2' :: (Ord a) => [a] -> a
maximum2' []     = error "maximum of empty list"
maximum2' [x]    = x
maximum2' (x:xs) = x `max` maximum2' xs
-- Se obtiene el mayor elemento de una lista pero ocupando la funcion "max"