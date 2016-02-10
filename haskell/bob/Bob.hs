module Bob (responseFor) where

import Data.Char
-- problem: doesn't recognize Unicode space characters
import Data.String.Utils (strip)
import Text.Regex (mkRegex, subRegex)

isYelling :: String -> Bool
isYelling s = (s == map toUpper s) && (s /= map toLower s)

isQuestion :: String -> Bool
isQuestion s = '?' == last s

-- Borrowed from http://exercism.io/submissions/2f39f9c9d6f3110cf0961c19
-- I tried figuring out how to create a Unicode regex that would match
-- Unicode space characters, but gave up.
isWhitespace :: String -> Bool
isWhitespace =
  all (`elem` "\n\r \t\v\xA0\x2002")

responseFor :: String -> String
responseFor s
  | isWhitespace s = "Fine. Be that way!"
  | isYelling s    = "Whoa, chill out!"
  | isQuestion s   = "Sure."
  | otherwise      = "Whatever."