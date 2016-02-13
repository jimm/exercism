module DNA (count, nucleotideCounts) where

import Data.Map (Map, fromList)

legit :: Char -> Bool
legit 'A' = True
legit 'C' = True
legit 'G' = True
legit 'T' = True
legit  c  = error $ "invalid nucleotide " ++ show c

count :: Char -> String -> Int
count c s
  | legit c && all id (map legit s) =
    length $ filter id (map (\ch -> c == ch) s)

nucleotideCounts :: String -> Map Char Int
nucleotideCounts s
  | all id (map legit s) = fromList $ map (\c -> (c, count c s)) "ACGT"
