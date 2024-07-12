cylinder :: (RealFloat a) => a -> a -> a
cylinder r h =
    let sideArea = 2 * pi * r * h
        topArea = pi * r *2 -- Se definen funciones ocupando la palabra reservada "let"
    in sideArea + 2 * topArea -- Se ocupan las funciones declaradas en "let" a traves de la palabra reservada "in"

-- La principal diferencia entre "let" y "where" es que "let" define expresiones mientras que "where" es unicamente una construccion sintactica


pruebaLet1 = 4 * (let a = 9 in a + 1) + 2
-- Se define una funcion que contiene una expresion "let" donde se define el valor de una variable y se utiliza en la expresion "in"

pruebaLet2 = [let square x = x * x in (square 1, 3, 5)]
-- Se define una funcion y se implmenta sobre una lista de tuplas

pruebaLet3 = (let a = 100; b = 200; c = 300 in a * b * c, let foo = "Hey"; boo = "there" in foo ++ boo)
-- Se definen 3 variables diferentes a traves del operador ";" en una misma linea