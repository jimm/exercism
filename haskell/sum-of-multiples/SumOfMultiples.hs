module SumOfMultiples (sumOfMultiples, sumOfMultiplesDefault) where
import Debug.Trace

sumOfMultiples :: [Int] -> Int -> Int
sumOfMultiples mults upperBound = sum $ map (sum $ multiplesOf upperBound [] 1) mults

sumOfMultiplesDefault :: Int -> Int
sumOfMultiplesDefault = sumOfMultiples [3, 5]

multiplesOf :: Int -> [Int] -> Int -> Int -> [Int]
multiplesOf upperBound acc n x
  | (x * n) >= upperBound = acc
  | otherwise             = multiplesOf x (n+1) upperBound (x*n : acc)
