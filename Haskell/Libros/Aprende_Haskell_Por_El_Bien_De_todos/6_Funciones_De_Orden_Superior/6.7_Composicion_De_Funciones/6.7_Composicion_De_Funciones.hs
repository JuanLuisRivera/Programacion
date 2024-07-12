a = map (\x -> negate (abs x)) [5,-3,-6,7,-3,2,-19,24] -- Funcion que convierte todos los numeros de una lista en negativos ocupando una funcion lambda
b = map (negate . abs) [5,-3,-6,7,-3,2,-19,24] -- Se define la misma funcion anterior pero ocupando la composicion de funciones

c = sum . replicate 5 . max 6.7 $ 8.9 -- Se aplica la funcion parcialmente aplicada "max" a los argumentos "6.7" y "8.9", de manera que regresa el "8.9", posteriormente
-- Aplica la funcion parcialmente aplicada "replicate 5" a "8.9", generando la lista "[8.9, 8.9, 8.9, 8.9, 8.9]" y posteriormente realiza la suma de todos los elemntos
-- En la lista, obteniendo como resultado final el valor "44.5"

sum' :: (Num a) => [a] -> a
sum' = foldl (+) 0 -- Se define la funcion "sum'" en "estilo libre de puntos", es decir no se especifica explicitamente los valores sobre los que actua, sino que al
-- Usarse "foldl (+) 0" se establece que se aplica sobre una lista

fn = ceiling . negate . tan . cos . max 50 -- Se define la funcion fn en "estilo libre de puntos"

oddSquareSum :: Integer
oddSquareSum = sum . takeWhile (<10000) . filter odd . map (^2) $ [1..] -- Se reescribe la funcion ocupando composicion de funciones

oddSquareSum :: Integer 
oddSquareSum =
    let oddSquares = filter odd $ map (^2) [1..]
        belowLimit = takeWhile (<10000) oddSquares
    in  sum belowLimit -- Para codigos complejos no se recomienda el uso de la composicion de funciones, en su lugar se recomienda la definicion de funciones
-- Intermedias de manera que se facilite la lectura y comprension del codigo