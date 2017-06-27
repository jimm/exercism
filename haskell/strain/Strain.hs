module Strain (keep, discard) where

keep :: (a -> Bool) -> [a] -> [a]
keep f xs = do_keep f xs []

-- TODO use foldr

do_keep :: (a -> Bool) -> [a] -> [a] -> [a]
do_keep _ [] acc = reverse acc
do_keep f (x:xs) acc
  | f x       = do_keep f xs (x:acc)
  | otherwise = do_keep f xs acc

discard :: (a -> Bool) -> [a] -> [a]
discard f = keep (not . f)
