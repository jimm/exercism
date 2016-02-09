module Grains (square, total) where
import Data.Bits (shift)

square :: Int -> Integer
square n = shift 1 (n-1)

total :: Integer
total = square 65 - 1
