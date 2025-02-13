def is_leap(year):
    leap = False #Se regresara False en todos los casos excepto en los que se cumpla la condicion de abajo
    
    if year % 4 == 0: # Si el año es divisible por 4
        if year % 100 == 0: # Si el año es divisible por 100
            if year % 400 == 0: # Si el año es divisible por 400
                leap = True
        else: #Si no es divisible por 100, es bisiesto
            leap = True
    
    return leap

year = int(input())
print(is_leap(year))