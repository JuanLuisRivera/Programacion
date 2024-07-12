--La palabra reservada 'let' se usa para definir un elemento en tiempo de ejecucion, 
--En este caso una lista, no se declara aqui porque no es necesario

--Los siguientes ejemplos muestran el funcionamiento de la concatenacion de listas

lostNumbers = [4,8,15,16,23,42] --Se define una lista de enteros y se inicializa con 6 elementos

a = [1,2,3,4,5] --Se define una lista de enteros y se inicializa con 5 elementos
b = [6,7,8,9,10] --Se define una lista de enteros y se inicializa con 5 elementos

c = a ++ b --Se concatena la lista a con la lista b a traves del operador "++" y se asigna a la variable c

list1 = ['h', 'o']

list2 = ['l','a']

hola = list1 ++ list2 --Se pueden concatenar listas de caracteres

--Es importante notar que la concatenacion con el operador "++" se realiza sobre el argumento izquierdo, de manera que debe recorrer
--La lista izquierda en completitud antes de agregar los elementos de la lista derecha, haciendo el proceso muy tardado sobre listas grandes

a1 = 1

b1 = [6,7,8,9,10]

c1 = a1: b1 --En contraste el operador ":" se usa para agregar elementos a una lista, de manera que "contatena" en el primer elemento
--De la lista, evitando la necesidad de recorrer en su totalidad la lista izquiera y por tanto, realiza el proceso de forma constante

letra = "Steve" !! 3 --Una cadena se considera una lista de caracteres y por tanto se pueden aplicar los operadores anteriores a las cadenas
--De fomrma que accederemos al tercer elemento de la cadena, donde el primer elemento es el 0 y corresponde a la "S"

list3 = [[], [1], [1,2,3]] --Se pueden tener listas de listas, las cuales deben respetar los mismos criterios de para su creacion que las
--Listas normales

--Los siguientes ejemplos muestran el funcionamiento de algunas funciones de listas
e = head [1,2,3] --Se accede al primer elemento de una lista, de manera que se obtiene el 1
f = tail [1,2,3] --Se accede a todos los elementos excepto el primero

h = last [1,2,3] --Se accede al ultimo elemento de una lista, de manera que se obtiene el 3
i = init [1,2,3] --Se accede a todos los elementos excepto el ultimo

--Es importante notar que estas funciones no pueden ser usadas en listas vacias ya que retornan un error

j = length [1,2,3] --Se obtiene la longitud de una lista

k = null [1,2,3] --Se indica si una lista es vacia o no, en caso de ser vacia se obtiene True, de lo contrario False

l = reverse [1,2,3] --Se invierten los elementos de una lista

m = take 2 [1,2,3,4,5] --Se extraen los primeros dos elementos de una lista

n = drop 2 [1,2,3,4,5] --Se quitan los primeros dos elementos de una lista

o = maximum [1,2,3,4,5] --Se obtiene el elemento maximo de una lista
 --Es importante notar que maximum se puede aplicar a cualquier lista de elemntos que se puedan ordenar

p = minimum [1,2,3,4,5] --Se obtiene el elemento minimo de una lista

q = sum [1,2,3,4,5] --Se obtiene la suma de todos los elementos de una lista

r = product [1,2,3,4,5] --Se obtiene el producto de todos los elementos de una lista

s = elem 3 [1,2,3,4,5] --Se indica si un elemento se encuentra en una lista

t = 3 `elem` [1,2,3,4,5] --Se indica si un elemento se encuentra en una lista con el operador de manera infija