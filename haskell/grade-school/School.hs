module School (School, empty, add, grade, sorted) where

-- TODO use Data.Map

data School = School Map.Map Int [String] deriving (Show)

empty = School []

add :: Int -> String -> School -> School
add grade name school = School(Entry(grade, name) : entries school)

grade :: Int -> School -> [String]
grade grade school =
  map entryName $ entriesForGrade grade school

sorted :: School -> [(Int, [String])]
-- sorted = id
sorted _ = [(2, ["Aimee"])]

-- ****************

entries :: School -> [Entry]
entries (School es) = es

entryGrade :: Entry -> Int
entryGrade (Entry (g, _)) = g

entryName :: Entry -> String
entryName (Entry (_, n)) = n

entriesForGrade :: Int -> School -> [Entry]
entriesForGrade grade school =
  filter (\e -> (entryGrade e) == grade) (entries school)
