a = sum (map sqrt [1..130]) -- Se establece la suma obligando a que se aplique la funcion "sqrt" a traves de "map" a todos los elementos primero y posteriomente se realice la suma
b = sum $ map sqrt [1..130] -- Se usa el operador "$" en el mismo calculo anterior, eliminando el uso de los parentesis ya que el operador tiene el valor de precedencia mas bajo
-- De manera que todas las funciones que le precedan se aplicaran al final

c = map ($ 3) [(4+), (10*), (^2), sqrt] -- Esta ejecucion nos regresaria un error si eliminamos el operador "$" debido a que la funcion "map" necesita una funcion para aplicar a
-- La lista