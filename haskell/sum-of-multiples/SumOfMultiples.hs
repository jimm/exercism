module SumOfMultiples (sumOfMultiples, sumOfMultiplesDefault) where

import Data.List (nub)

sumOfMultiples :: [Int] -> Int -> Int
sumOfMultiples mults upperBound = sum $ nub ( concatMap (multsOfSingle upperBound) mults )

sumOfMultiplesDefault :: Int -> Int
sumOfMultiplesDefault = sumOfMultiples [3, 5]

multsOfSingle :: Int -> Int -> [Int]
multsOfSingle upperBound x = multiplesOf upperBound [] 1 x

multiplesOf :: Int -> [Int] -> Int -> Int -> [Int]
multiplesOf upperBound acc n x
  | (x * n) >= upperBound = acc
  | otherwise             = multiplesOf upperBound (x*n : acc) (n+1) x
