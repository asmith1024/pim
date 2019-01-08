import Test.QuickCheck
import Poly

prop_x1 x = polY x [1.0, 2.0, 3.0] == 1.0 + (x * 2.0) + (x**2.0 * 3.0)

prop_x2 x = polY x [1.0, -2.0, 3.0, -4.0] == 1.0 + (x * (-2.0)) + (x**2.0 * 3.0) + (x**3.0 * (-4.0))

prop_x3 x = polY x [1.0, 0.0, 0.0, 0.0, 5.0] == 1.0 + (x**4.0 * 5.0)