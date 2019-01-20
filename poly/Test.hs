import Test.QuickCheck
import Poly

-- from ghci...
-- :load Test.hs
-- quickCheck prop_name

prop_x1 :: Double -> Bool
prop_x1 x = polY x [1.0, 2.0, 3.0] == 1.0 + (x * 2.0) + (x^^2 * 3.0)

prop_x2 :: Double -> Bool
prop_x2 x = polY x [1.0, -2.0, 3.0, -4.0] == 1.0 + (x * (-2.0)) + (x^^2 * 3.0) + (x^^3 * (-4.0))

prop_x3 :: Double -> Bool
prop_x3 x = polY x [1.0, 0.0, 0.0, 0.0, 5.0] == 1.0 + (x^^4 * 5.0)

-- how to use QuickCheck to test printing?

