module DNA (hammingDistance) where

hammingDistance :: [Char] -> [Char] -> Int
hammingDistance s1 s2 = hd 0 s1 s2

hd :: Int -> [Char] -> [Char] -> Int
hd acc [] [] = acc
hd acc (c1:s1) (c2:s2)
  | c1 == c2  = hd acc     s1 s2 
  | otherwise = hd (acc+1) s1 s2 
