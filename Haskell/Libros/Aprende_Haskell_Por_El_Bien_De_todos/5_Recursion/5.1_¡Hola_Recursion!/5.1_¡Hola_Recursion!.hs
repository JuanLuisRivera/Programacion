fibonacci :: (Integral a) => a -> a
fibonacci 0 = 1
fibonacci 1 = 1
fibonacci x = x * fibonacci (x - 1)
-- Se hace la definicion de una funcion de forma recursiva para plantear cual es el proposto de la recursividad