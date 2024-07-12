quicksort :: (Ord a) => [a] -> [a]
quicksort [] = []
quicksort (x: xs) = -- Se generan  dos listas a traves de dos funciones auxiliares
    let smallerSorted = quicksort [a | a <- xs, a <= x] -- Lista intensional donde los elementos son menores que el elemento "x"
        biggerSorted = quicksort [a | a <- xs, a >= x] -- Lista intensional donde los elementos son mayores que el elemento "x"
    in smallerSorted ++ [x] ++ biggerSorted -- Lista total de elementos

