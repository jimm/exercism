module LeapYear (isLeapYear) where

isLeapYear :: Int -> Bool
isLeapYear x
    | (x `rem` 400 == 0) = True
    | (x `rem` 100 == 0) = False
    | (x `rem`   4 == 0) = True
    | otherwise = False
