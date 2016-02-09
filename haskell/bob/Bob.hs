module Bob (responseFor) where

import Data.Char
-- problem: doesn't recognize Unicode space characters
import Data.String.Utils (strip)

isYelling :: String -> Bool
isYelling s = s == map toUpper s

isQuestion :: String -> Bool
isQuestion s = '?' == last s

isEmpty :: String -> Bool
isEmpty s = length (strip s) == 0

responseFor :: String -> String
responseFor s
    | isEmpty s    = "Fine. Be that way!"
    | isYelling s  = "Whoa, chill out!"
    | isQuestion s = "Sure."
    | otherwise    = "Whatever."