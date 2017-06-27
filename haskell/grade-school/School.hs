module School (School, empty, add, grade, sorted) where

import qualified Data.Map.Strict as Map
import Data.List (sort)

data School a b = School a [b] deriving (Show)

empty :: School a [b]
empty = School(Map.empty)

add :: Int -> String -> School a [b] -> School a [b]
add grade name school =
  School(Map.insert grade (name:students))
  where students = Map.findWithDefault school grade []

grade :: Int -> School a [b] -> [String]
grade grade school =
  Map.lookup school grade

sorted :: School Int [String] -> [(Int, [String])]
-- sorted = id
sorted school =
  sort foldr (\names acc -> acc ++ names) [] school
