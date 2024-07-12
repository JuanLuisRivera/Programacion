a = max 4 5 -- Se ejemplifica como es que la funcion se ejecuta tomando dos elementos "4" y "5"

b = (max 4) 5 -- Se ejemplifica como es que la funcion trabaja tomando unicamente el valor "4" y posteriormente el valor "5"

c = a == b -- Se verifica que ambos resultados son identicos

-- La diferencia entre ambas aplicaciones es que en el primer caso la funcion se aplica sobre dos argumentos, mientras que en la segunda
-- Se aplica unicamente sobre un argumento y posteriomente se vuelve a aplicar sobre el segundo argumento, es esa propiedad de ser
-- Una aplicacion incompleta de la funcion "max" se conoce como "aplicacion parcial" y tiene bastante relevancia porque es posible ocupar
-- Funciones parcialmente aplicadas como argumentos de otras funciones

multThree :: (Num a) => a -> a -> a -> a -- Definimos una funcion que toma tres elementos
multThree x y z = x * y * z

d = multThree 9 -- Se ejecuta la funcion con unicamente un elemento, conviertiendola en una funcion parcialmente aplicada
-- Esta funcion no se puede mostrar ya que no pertenece a la clase "Show" y por tanto, llamarla nos generaria un error
-- Sin embargo la ejecucion misma de la funcion si se ejecuta sin ningun problema
e = d 2 4 -- Se termina de aplicar los demas argumentos de la funcion, obteniendo un valor entero como resultado
f = d 6 3

divideByTen :: (Floating a) => a -> a
divideByTen = (/10) -- se usa una funcion infija para ser aplicada parcialmente

g = divideByTen 200 --Se ocupa la funcion parcialmente aplicada con el valor numerico 200