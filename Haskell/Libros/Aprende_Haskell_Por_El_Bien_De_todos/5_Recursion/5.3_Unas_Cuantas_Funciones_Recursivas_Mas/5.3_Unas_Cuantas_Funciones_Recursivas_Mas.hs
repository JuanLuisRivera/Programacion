replicate' :: (Num i, Ord i) => i -> a -> [a] -- Se define la funcion donde es necesario que la variable "i" sea
-- Ordenable ya que nos interesa poder compararla con el "0" y debe ser un tipo "Num" para poder operar matematicamente con ella
replicate' x y -- X sera la variable que indicara la cantidad de veces que la variable se va a repetir, tendra las propiedades de "i"
    | x <= 0 = [] --Si la variable x es menor o igual a 0 se retorna una lista vacia
    | otherwise = y: replicate'(x - 1) y -- Se agrega la variable "y" a la lista de manera recursiva

take' :: (Ord i, Num i) => i -> [a] -> [a]
take' x _
    | x <= 0 = [] -- Caso donde el numero de elementos que se quieren tomar sera menor 
take' _ [] = [] -- Caso donde se toman elementos de una lista vacia
take' i (x: xs) = x: take' (i - 1) xs -- Se tomara un elemento de la lista y se aplicara recursion sobre los demas elementos hasta llegar
-- A los casos base

reverse' :: [a] -> [a]
reverse' [] = [] -- Caso base
reverse' (x: xs) = reverse' xs ++ [x] -- El operador "++" concatena el elemento "[x]" al final de la lista

repeat' :: a -> [a]
repeat' x = x: repeat' x -- Se genera una lista infinita de elementos "x"

zip' :: [a] -> [b] -> [(a, b)]
zip' _ [] = []
zip' [] _ = []
zip' (x:xs) (y:ys) = (x, y): zip' xs ys

elem' :: (Eq a) => a -> [a] -> Bool -- En este caso declaramos que nos interesa que el elemnto pueda ser comparado, por lo cual
-- Declaramos la clase de tipo del elemento "a" a "Eq"
elem' _ [] = False
elem' x (y: ys)
    | x == y = True
    | x /= y = False