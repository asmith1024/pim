import Data.Word (Word16)
import System.Random (StdGen, mkStdGen)
import System.Random.Shuffle (shuffle')
import Test.QuickCheck
import Sets

-- define a custom generator for Prefs
-- generating pseudo-random preference sets is not trivial in Haskell - see https://wiki.haskell.org/Random_shuffle
-- I hate dependencies.

-- let's try to get QuickCheck to handle that for us:

initPreferences :: Word16 -> ([[Int]], [[Int]])
initPreferences size = let g = mkStdGen (fromIntegral size :: Int)
                        in (preflist size g, preflist size g)

preflist :: Word16 -> StdGen -> [[Int]]
preflist size g = let c = fromIntegral size :: Int
                   in [prefs c g | x <- [1..c]]

prefs :: Int -> StdGen -> [Int]
prefs size = shuffle' [1 .. size] size

-- from ghci...
-- :load Test.hs
-- quickCheck prop_name

prop_t1 :: Word16 -> Bool
prop_t1 size = let (a, b) = initPreferences size
                in areAllocationsStable a b (allocatePreferences a b) 
