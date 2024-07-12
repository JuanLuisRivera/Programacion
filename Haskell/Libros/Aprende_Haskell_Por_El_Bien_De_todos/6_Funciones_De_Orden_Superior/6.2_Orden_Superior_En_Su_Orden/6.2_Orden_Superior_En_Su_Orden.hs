applyTwice :: (a -> a) -> a -> a -- Se define una funcion que toma una funcion y la aplica dos veces a un elemento
-- Donde la primera parte de la declaracion especifica (entre los parentesis) que sera una funcion con un argumento "a" 
-- Y el segundo argumento sera un elemento "a", y se regresara un elemento del tipo "a"
applyTwice f x = f (f x)

a = applyTwice (+ 3) 10 -- Se aplican dos veces la funcion "+", de manera que se obtiene "6" y el operador +, por tanto
-- Se suma a 10 al resultado obtenido por 6, dando 16

b = applyTwice (++ " HAHA") "HEY" -- De igual manera se aplica dos veces la funcion "++" sobre las cadenas y al final se regresa la funcion 
-- "++ HAHA HAHA HEY" donde como "++" es una funcion parcialmente aplicada despues de aplicarse la funcion "applyTwice" se vuelve a aplicar
-- La cadena "HEY"

zipWith' :: (a -> b -> c) -> [a] -> [b] -> [c] -- Se toma una funcion con dos elementos, posteriormente se aplica sobre 2 listas "[a]" y "[b]"  
-- Y da como resultado una tercera lista "[c]"
zipWith' _ [] _ = [] 
zipWith' _ _ [] = []
zipWith' f (x:xs) (y:ys) = f x y : zipWith' f xs ys -- Se aplica la funcion "f" sobre los dos elementos "x" y "y", los cuales se agregan a una
-- Lista de forma recursiva sobre los demas elementos de las listas

c = zipWith' (+) [1, 2, 3] [1, 2, 3] -- Se aplicara la funcion "+" sobre los elementos de la funcion de manera recursiva, dando como resultado
-- Una lista de la suma de todos los elementos de una lista con los elementos de la otra lista, es decir : [2, 4, 6]

d = zipWith' (++) ["Hola ", "Adios "] ["Mundo", "Comida"] -- De manera similar, el resultado sera una lista ["Hola Mundo", "Adios Comida"]

flip' :: (a -> b -> c) -> (b -> a -> c) -- Se define la funcion que toma una funcion y regresa otra funcion con sus parametros intercambiados
flip' f = g -- "f" es una funcion que toma como argumentos de tipo "a" y "b" produciendo una salida de tipo "c", 
-- Mientras que "g" es una funcion que toma como argumentos de tipo "b" y "a" y produce una salida de tipo "c", cambaindo por tanto de posicion
-- Los argumentos de la funcion original conservando la misma funcion al igualar "f" y "g"
    where g x y = f y x -- Es necesario declarar que "f" es igual a "g" y por tanto se necesita ocupar "where" para asegurar la igualdad
    -- Entre las funciones

flip'' :: (a -> b -> c) -> (b -> a -> c) -- Otra forma de declarar la funcion flip' sin el uso de la expresion "where"
flip'' f x y = f y x

e = flip' zip [1,2,3,4,5] "hello" -- Se intercambian los argumentos de la funcion, de forma que, sin la funcion "flip'" se esperaria:
-- [(1, "h"), (2, "e"), (3, "l"), ...], el resultado con "flip'" es : [("h", 1), ("e", 2), ("l", 3), ...]