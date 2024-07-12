-- Los modulos son colecciones de funciones tipos y clases de tipos que estan relacionadas entre si, de forma que permite la reutilizacion del mismo codigo en diferentes programas.
-- Un programa en haskell es un conjunto de modulos donde un modulo principal carga otros modulos y con ellos realiza una tarea especifica

import Data.List -- Se importa un modulo a traves de la palabra reservada "import" seguido del nombre del modulo
-- import Data.List (nub, sort) Importa solo las funciones "nub" y "sort" del modulo "Data.List"
-- import Data.List hiding (nub) Importa todas las funciones del modulo "Data.List" a excepcion de la funcion "nub"
import qualified Data.Map -- Si se importan modulos donde existen funciones que conflictuan en nombre con las de otros modulos, se deberan ocupar las funciones haciendo referencia a la libreria,
-- Por ejemplo: Data.Map.filter llama a la funcion "filter" del modulo "Data.Map" y no a la funcion "filter" estandar

-- import qualified Data.Map as M -- Renombramos la forma de llamar al modulo de manera que no es necesario escribir "Data.Map", sino que es posible hacerlo unicamente con "M", por ejemplo: "M.filter"
-- Llama a la funcion "filter" del modulo "Data.Map"

numUniques :: (Eq a) => [a] -> Int -- Funcion que nos dice el numero de elementos unicos en una lista
numUniques = length . nub -- La funcion "nub" se importa del modulo "Data.list", la cual toma una lista y regresa otra sin duplicados

-- Si se desean exportar modulos al trabajar en "ghci" se usa la siguiente sintaxis :m + "modulo", por ejemplo: ":m + Data.List"; si se desean importar varios modulos: :m + Data.List Data.Map Data.Set

