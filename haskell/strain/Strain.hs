module Strain (keep, discard) where

keep :: (a -> Bool) -> [a] -> [a]
keep f xs = doKeep f xs []

doKeep :: (a -> Bool) -> [a] -> [a] -> [a]
doKeep _ [] acc = reverse acc
doKeep f (x:xs) acc
  | f(x)      = doKeep f xs (x:acc)
  | otherwise = doKeep f xs acc

discard :: (a -> Bool) -> [a] -> [a]
discard f = keep (not . f)
