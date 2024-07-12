a = (8, 11)

b = fst a --Este operador toma unicamente el primer elemento de La dupla "a" y lo asigna a la variable "b"
c = snd a --Este operador toma el valor del segundo elemento de la dupla "a" y lo asigna a la variable "c"

--Es importante notar que estos operadores funcionan unicamente con tuplas de dos elementos, tambien llamadas "duplas"

d = [1, 2]
e = ["Hola", "Bola"]
f = zip d e --Se utiliza la funcion "zip" para generar una lista de duplas, donde los elementos se relacionaran uno a uno de acuerdo al orden de aparicion
-- En su respectiva lista

triangles1 = [ (a, b, c) | c <- [1..10], b <- [1..10], a <- [1..10]] --Se genera una lista de triplas, producto de la combinacion de los diferentes
--Numeros del 1 al 10

triangles2 = [ (a, b, c) | c <- [1..10], b <- [1..c], a <- [1..b], a^2 + b^2 == c^2] --Garantizamos que las triplas unicamente seran aquellas que
--Satisfacen el teorema de pitagoras y al mismo tiempo satisfagan las condiciones de no ser mayores los lados que la hipotenusa y que no seran mas 
--Largos los lados uno respecto al otro, de manera que garantizamos que los lados seran triangulos rectangulos

triangles3 = [ (a, b, c) | c <- [1..10], b <- [1..c], a <- [1..b], a^2 + b^2 == c^2, a + b + c == 24] --Garantizamos que los unicos triangulos en la lista
--Seran aquellos que cumplan la propiedad de tener un perimetro igual a 24