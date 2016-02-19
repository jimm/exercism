module Gigasecond (fromDay) where

import Data.Time.Clock (UTCTime, NominalDiffTime, addUTCTime)

gigaseconds :: NominalDiffTime
gigaseconds = 1e9

fromDay :: UTCTime -> UTCTime
fromDay = addUTCTime gigaseconds
