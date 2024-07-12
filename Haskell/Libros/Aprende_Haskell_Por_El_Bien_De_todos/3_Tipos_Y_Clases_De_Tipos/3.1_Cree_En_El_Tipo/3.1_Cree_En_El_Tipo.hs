--Haskell es un sistema de tipos estatico que ocupa un sistema de inferencia de tipos
--Para conocer el tipo de dato de las expresiones en el proceso de compilacion.
--Si se desea conocer el tipo de una expresion en tiempo de ejecucion se ocupa el comando ":t"
--Y posteriormente la expresion

--De manera similar, al definir una funcion se ocupan "->" para separar los argumentos 
--Y para separar tambien tipo que se retorna


funcionMuestraInt :: Int -> Int -> Int --Se define una funcion "funcionMuestraInt" que recibira dos enteros
--Y retornara un entero
funcionMuestraInt x y = x + y
--El tipo de dato "Int" esta acotado entre -2147483648 y 2147483647

funcionMuestraInteger :: Integer -> Integer -> Integer
funcionMuestraInteger x y = x + y
--El tipo de dato "Integer" no es acotado, por lo que puede representar valores numericos muy grandes
--Pero es menos eficiente que el tipo de dato "Int"

funcionMuestraFloat :: Float -> Float -> Float
funcionMuestraFloat x y = x + y
--El tipo de dato "Float" tiene precion simple

funcionMuestraDouble :: Double -> Double -> Double
funcionMuestraDouble x y = x + y
--El tipo de dato "Double" tiene doble precision

funcionMuestraBoolean :: Bool -> Bool -> Bool
funcionMuestraBoolean x y = x || y
--El tipo de dato "Bool" solo puede ser True o False

funcionMuestraChar :: Char -> Char
funcionMuestraChar x = 'x' 
--La funcion regresa un elemento de tipo char