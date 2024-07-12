a = [x * 2 | x <- [1..10]] --Se define una lista intensional donde x es el doble de cada elemento de la lista [1..10]

b = [x*2 | x <- [1..10], x*2 >= 12] --Se define una lista intensional la cual esta acotada por la condicion de que el
--Producto debe ser mayor o igual a 12

c = [x | x <- [50 .. 100], x `mod` 7 == 3] --Se define una lista intensional la cual contiene a todos los elementos
--Que cumplen la condicion de modulo 7 = 3

boomBang xs = [if x < 10 then "BOOM" else "BANG" | x <- xs, odd x] --Se define la funcion "boomBang" la cual evalua
--Los elementos de la lista "xs", de manera que si el elemento es menor que 10 e impar, se cambia el valor por la
--Cadena "BOOM" y si es mayor a 10 e impar, el valor se cambiara por "BANG", eliminando aquellos numero que no sean
--Impares de la lista

notOfTheList = [x | x <- [1..20], x /= 13, x /= 15, x /= 19] --Se define una lista por comprension, la cual se filtra
--Estableciendo que si los elementos son diferentes a 13, 15 o 19, perteneceran a la lista

multiplyLists a b = [x * y | x <- a, y <- b] --Se define la funcion "multiplyLists" la cual genera el producto de los
--Elementos de dos listas

multiplyListsOVer50 a b = [x * y | x <- a, y <- b, x * y > 50] --Se define la funcion "multiplyListsOVer50" que toma
--Los elementos de dos listas y regresa una lista intensional si el resultado de la multiplicacion es mayor a 50

nombres = ["Lupe", "Okani", "Luis"]
adjetivos = ["Graciosa", "Rica", "Curioso"]

nombresAdjetivos = [nombre ++ " " ++ adjetivo | nombre <- nombres, adjetivo <- adjetivos] --Se define una lista
--Intensional la cual combina todos los nombres con todos los adjetivos de ambas listas

length' xs = sum [1 | _ <- xs] --Se define la funcion "length'" la cual cuenta la cantidad de elementos de una lista
--Sin importar el tipo de dato de la lista. El operador "_" se refiere a que no importa lo que se extraiga de la lista
--Lo reemplazara por 1 y despues de reemplazar todos los elementos, los sumara para obtener la cantidad total

removeNonUppercase st = [c | c <- st, c `elem` ['A' .. 'Z']] --Debido a que lasn cadenas son listas de caracteres
--Se pueden ocupar listas intensionales para procesar y producir nuevas cadenas

xxs = [[1,3,5,2,3,1,2,4,5],[1,2,3,4,5,6,7,8,9],[1,2,4,2,1,6,3,1,3,2,3,6]] --Se define la lista de listas "xxs"
ys = [[ x | x <- xs, even x ] | xs <- xxs] --Se define una lista "ys" a partir de "xxs", la cual generara una lista de listas que cumpliran la condicion
--De que los elementos de las listas dentro de "xxs" sean pares