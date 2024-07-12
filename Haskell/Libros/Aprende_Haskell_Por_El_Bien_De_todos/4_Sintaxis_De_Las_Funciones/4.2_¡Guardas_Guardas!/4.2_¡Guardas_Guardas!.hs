ejemplosGuarda :: (Integral a) => a -> String
ejemplosGuarda x
    | x <= 18 = "Hola"
    | x <= 22 ="Bola"
    | x <= 35 ="Pelota"
    | otherwise = "Nada"
--Se ocupan los llamados guardas "|" los cuales son operadores que se asemejan al ajuste de patrones pero toman una variable y
--Determinan si dicha variable satisface una condicion booleana establecida y de la misma manera se establece un caso general
--a traves de la palabra "otherwise"

max' :: (Ord a) => a -> a -> a
max' a b
    | a > b = a
    | otherwise = b
--Esta funcion toma dos elementos los cuales pueden ser comparados y regresan el elemento que es mayor