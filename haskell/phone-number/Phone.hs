module Phone (areaCode, number, prettyPrint) where

import Data.Char (isDigit)
import Text.Printf (printf)

onlyDigits :: String -> String
onlyDigits s = [ ch | ch <- s, isDigit ch ]

isValidWithLeading1 :: String -> Bool
isValidWithLeading1 s =
  length ds == 11 && first == '1'
  where ds = onlyDigits(s)
        first = head (take 1 s)

isValid :: String -> Bool
isValid s =
  length ds == 10 || (first == '1' && (length rest) == 10)
  where ds = onlyDigits(s)
        first = head (take 1 s)
        rest = drop 1 s

number :: String -> String
number s
  | length s == 0 = "0000000000"
  | isValidWithLeading1 s = drop 1 $ onlyDigits s
  | isValid s = onlyDigits s
  | otherwise = "0000000000"

areaCode :: String -> String
areaCode s = take 3 (number s)

firstThree :: String -> String
firstThree s = take 3 $ drop 3 (number s)

lastFour :: String -> String
lastFour s = take 4 $ drop 6 (number s)

prettyPrint :: String -> String
prettyPrint s =
  -- I can't seem to get this printf to work. Sigh.
  -- printf "(%d) %d-%d" (areaCode s) (firstThree s) (lastFour s)
  "(" ++ areaCode s ++ ") " ++ firstThree s ++ "-" ++ lastFour s
