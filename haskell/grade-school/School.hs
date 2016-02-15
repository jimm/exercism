module School (empty, grade, sorted) where

data Grade = Grade Int deriving (Show)
data Name = Name String deriving (Show)
data Entry = Entry (Grade, Name) deriving (Show)
data School = School [Entry] deriving (Show)

empty :: School
empty = School []

add :: Int -> String -> School
add grade name = empty

grade :: Int -> School -> [String]
grade grade school =
  map snd $ filter (\Entry(g, n) -> g == grade) school

sorted :: School -> School
sorted = id
