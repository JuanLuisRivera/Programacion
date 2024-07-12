doubleMe x = x + x --Define la funcion doubleMe la cual toma un numero y regresa el doble de ese numero

doubleUs x y = x * 2 + y * 2 --Define la funcion doubleUs la cual toma dos numeros y regresa el doble de ambos

doubleUs' x y = (doubleMe x) + (doubleMe y) --Define la funcion doubleUs' la cual toma dos numeros y regresa el doble de ambos
-- pero usando la función doubleMe

doubleSmallNumber x = if x > 100 then x else x*2 --Define la funcion doubleSmallNumber la cual toma un numero y regresa el
-- doble de ese numero si el numero es menor a 100 y en caso contrario regresa el mismo numero

doubleSmallNumber' x = (doubleSmallNumber x) + 1 --Define la funcion doubleSmallNumber' la cual toma 
-- suma un 1 al resultado de la función doubleSmallNumber

miNombre = "Soy yo, Juan Luis" --Define una funcion la cual regresa 

--Es importante notar que las funciones no pueden empezar por mayuscula y que existen funciones que no toman ningun argumento 
