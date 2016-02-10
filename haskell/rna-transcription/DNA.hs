module DNA (toRNA) where

mapCGAT :: Char -> Char
mapCGAT 'C' = 'G'
mapCGAT 'G' = 'C'
mapCGAT 'A' = 'U'
mapCGAT 'T' = 'A'
mapCGAT  _  = ' '

toRNA :: String -> String
toRNA = map mapCGAT
