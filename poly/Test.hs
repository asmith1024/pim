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

-- beyond here we are not using QuickCheck to generate test data for us

prop_poly0 :: Bool
prop_poly0 = interpolatePoly [] == []

prop_polydup :: Bool
prop_polydup = interpolatePoly [(1.0,6.0),(2.0,17.0),(3.0,34.0),(2.0,4.0)] == []

prop_poly1 :: Bool
prop_poly1 = interpolatePoly [(1.0,6.0),(2.0,17.0),(3.0,34.0)] == [1.0,2.0,3.0]

prop_poly2 :: Bool
prop_poly2 = interpolatePoly [(3.0,34.0),(1.0,6.0),(2.0,17.0)] == [1.0,2.0,3.0]

prop_poly3 :: Bool
prop_poly3 = map round (interpolatePoly [(-2.0,25.0),(1.0,16.0),(0.0,1.0),(3.0,1600.0),(-1.0,0.0),(-3.0,400.0),(2.0,225.0)]) == [1,2,3,4,3,2,1]

prop_print1 :: Bool
prop_print1 = printPoly [1.0,-2.0,3.0,-4.0] == "f(x) = -4.0x^3 + 3.0x^2 - 2.0x + 1.0"

prop_print2 :: Bool
prop_print2 = printPoly [1.0,0.0,0.0,0.0] == "f(x) = 1.0"

prop_print3 :: Bool
prop_print3 = printPoly [0.0,0.0,0.0,-8.123,0.0,0.0,0.0,0.0] == "f(x) = -8.123x^3"

prop_print4 :: Bool
prop_print4 = printPoly [-6.444,0.0,0.0,-8.123,0.0,0.0,0.0,0.0] == "f(x) = -8.123x^3 - 6.444"

prop_printempty :: Bool
prop_printempty = printPoly [] == ""

prop_printzeroes :: Bool
prop_printzeroes = printPoly [0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0] == ""