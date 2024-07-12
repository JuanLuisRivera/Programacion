ejemplosGuarda2 :: (Integral a) => a -> a -> String
ejemplosGuarda2 x y
    | x + y <= 18 = "Hola"
    | x + y <= 22 ="Bola"
    | x + y <= 35 ="Pelota"
    | otherwise = "Nada"
-- Repetimos 3 veces la operacion "x + y"

ejemplosGuarda3 :: (Integral a) => a -> a -> String
ejemplosGuarda3 x y
    | z <= 18 = "Hola"
    | z <= 22 ="Bola"
    | z <= 35 ="Pelota"
    | otherwise = "Nada"
    where z = x + y
-- Se evita la triple declaracion de la suma al crear una variable "monigote" o "dummy", la cual seria la variable "z" en este caso

ejemplosGuarda4 :: (Integral a) => a -> a -> String
ejemplosGuarda4 x y
    | z <= a = "Hola"
    | z <= b ="Bola"
    | z <= c ="Pelota"
    | otherwise = "Nada"
    where z = x + y
          a = 18
          b = 22
          c = 35
-- Declaramos mas variables monigote para los valores estaticos

calcBmis :: (RealFloat a) => [(a, a)] -> [a]
calcBmis xs = [bmi w h | (w, h) <- xs]
    where bmi weight height = weight / height ^ 2 -- Se define una funcion auxiliar dentro de la seccion "where"