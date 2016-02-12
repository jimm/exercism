module Anagram (anagramsFor) where

import Data.List (sort)
import Data.Char (toLower)

anagramsFor :: String -> [String] -> [String]
anagramsFor s = doAnagramsFor [] (map toLower s)

doAnagramsFor :: [String] -> String -> [String] -> [String]
doAnagramsFor acc _ [] = reverse acc
doAnagramsFor acc s (x:xs)
  | isAnagram s (map toLower x) = doAnagramsFor (x:acc) s xs
  | otherwise                   = doAnagramsFor acc s xs

isAnagram :: String -> String -> Bool
isAnagram s1 s2 = s1 /= s2 && (sort s1) == (sort s2)