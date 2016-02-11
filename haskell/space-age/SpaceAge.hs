module SpaceAge (Planet(..), ageOn) where

data Planet = Mercury | Venus | Earth | Mars | Saturn | Jupiter | Neptune | Uranus
  deriving (Show)

earthYearInSeconds :: Double
earthYearInSeconds = 365.25 * 24 * 60 * 60

ageOn :: Planet -> Integer -> Double
ageOn p secs =  (fromIntegral secs) / (earthYearInSeconds * (yearLength p))

yearLength Earth = 1.0
yearLength Mercury = 0.2408467
yearLength Venus = 0.61519726
yearLength Mars = 1.8808158
yearLength Jupiter = 11.862615
yearLength Saturn = 29.447498
yearLength Uranus = 84.016846
yearLength Neptune = 164.79132
