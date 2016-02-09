module Sublist (Sublist(Equal, Sublist, Superlist, Unequal), sublist) where

data Sublist = Equal | Sublist | Superlist | Unequal
  deriving (Show, Eq)

sublist :: (Eq a) => [a] -> [a] -> Sublist
sublist x y
    | x == y              = Equal
    | length x < length y = sublistCheck x y
    | length y < length x = superlistCheck x y
    | otherwise           = Unequal

sublistCheck :: (Eq a) => [a] -> [a] -> Sublist
sublistCheck x y
    | x == y                 = Sublist
    | length x > length y    = Unequal
    | x == take (length x) y = Sublist
    | otherwise              = sublistCheck x (tail y)

superlistCheck :: (Eq a) => [a] -> [a] -> Sublist
superlistCheck x y
    | x == y                 = Superlist
    | length x < length y    = Unequal
    | y == take (length y) x = Superlist
    | otherwise              = superlistCheck (tail x) y
