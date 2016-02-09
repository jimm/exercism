module DNA (toRNA) where

mapCGAT :: Char -> Char
mapCGAT c
  | c == 'C'  = 'G'
  | c == 'G'  = 'C'
  | c == 'A'  = 'U'
  | c == 'T'  = 'A'
  | otherwise = ' '

toRNA :: String -> String
toRNA s = map mapCGAT s
