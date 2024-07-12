headCase1 :: [a] -> a
headCase1 [] = error "¡head no funciona con listas vacías!"
headCase1 (x:_) = x
-- Se define una funcion a traves de ajuste de patrones

headCase2 :: [a] -> a
headCase2 xs = case xs of [] -> error "¡head no funciona con listas vacías!"
                          (x: _) -> x
-- Se define una funcion a traves del uso de la expresion "case"

describeList :: [a] -> String
describeList xs = "La lista es" ++ case xs of []  -> "una lista vacía."
                                              [x] -> "una lista unitaria."
                                              xs  -> "una lista larga."
-- Una de las ventajas de las expresiones "case" es que se pueden ocupar en casi cualquier lugar, a diferencia del ajuste de patrones
--El cual unicamente se puede utilizar al definir una funcion