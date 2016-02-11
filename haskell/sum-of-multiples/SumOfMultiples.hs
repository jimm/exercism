module SumOfMultiples (sumOfMultiples, sumOfMultiplesDefault) where

-- FIXME: does not de-dupe multipliers that are common between all
-- Data.List.nub is a `uniq` function, see also `concatMap`

sumOfMultiples :: [Int] -> Int -> Int
sumOfMultiples mults upperBound = sum $ map (multsOfSingle upperBound) mults

sumOfMultiplesDefault :: Int -> Int
sumOfMultiplesDefault = sumOfMultiples [3, 5]

multsOfSingle :: Int -> Int -> Int
multsOfSingle upperBound x = sum $ multiplesOf upperBound [] 1 x

multiplesOf :: Int -> [Int] -> Int -> Int -> [Int]
multiplesOf upperBound acc n x
  | (x * n) >= upperBound = acc
  | otherwise             = multiplesOf upperBound (x*n : acc) (n+1) x
