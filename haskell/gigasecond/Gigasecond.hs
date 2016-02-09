module Gigasecond (fromDay) where

import Data.Time.Clock (UTCTime, NominalDiffTime, addUTCTime)

gigaseconds :: NominalDiffTime
gigaseconds = 1000000000

fromDay :: UTCTime -> UTCTime
fromDay t = addUTCTime gigaseconds t
