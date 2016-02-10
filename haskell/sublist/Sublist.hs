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
sublistCheck x y = superSub Sublist x y

superlistCheck :: (Eq a) => [a] -> [a] -> Sublist
superlistCheck x y = superSub Superlist y x

superSub :: (Eq a) => Sublist -> [a] -> [a] -> Sublist
superSub slist x y
  | x == y                 = slist
  | length x > length y    = Unequal
  | x == take (length x) y = slist
  | otherwise              = superSub slist x (tail y)
