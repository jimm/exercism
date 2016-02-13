module DNA (hammingDistance) where

hammingDistance :: [Char] -> [Char] -> Int
hammingDistance s1 s2 = sum $ map diffToInt (zip s1 s2)

diffToInt :: (Char, Char) -> Int
diffToInt (c1, c2)
  | c1 == c2  = 0
  | otherwise = 1
