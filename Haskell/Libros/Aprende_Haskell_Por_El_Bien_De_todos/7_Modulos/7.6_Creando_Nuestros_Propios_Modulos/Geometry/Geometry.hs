module Geometry -- Definimos un modulo ocupando la palabra reservada "module", seguido del argumento cuyo nombre debe corresponder al archivo en el cual estara
-- Guardada la informacion relativa al modulo
( sphereVolume -- Especificamos las funciones que se van a exportar de este modulo
, sphereArea
, cubeVolume
, cubeArea
, cuboidArea
, cuboidVolume
) where

sphereVolume :: Float -> Float -- Se definen las funciones del modulo
sphereVolume radius = (4.0 / 3.0) * pi * (radius ^ 3)

sphereArea :: Float -> Float
sphereArea radius = 4 * pi * (radius ^ 2)

cubeVolume :: Float -> Float
cubeVolume side = cuboidVolume side side side

cubeArea :: Float -> Float
cubeArea side = cuboidArea side side side

cuboidVolume :: Float -> Float -> Float -> Float
cuboidVolume a b c = rectangleArea a b * c

cuboidArea :: Float -> Float -> Float -> Float
cuboidArea a b c = rectangleArea a b * 2 + rectangleArea a c * 2 + rectangleArea c b * 2

rectangleArea :: Float -> Float -> Float -- Se definen funciones "intermedias", las cuales se ocupan para realizar calculos dentro del modulo, pero al no declararse como funciones a exportarse, no se pueden ocupar fuera del modulo.
rectangleArea a b = a * b