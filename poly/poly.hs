module Poly (evalpoly) where

import Data.List (foldl')

evalpoly :: (Eq a, Floating a) => [a] -> a -> a
evalpoly cs x = snd (evaltopair cs x) 

evaltopair :: (Eq a, Floating a) => [a] -> a -> (a, a)
evaltopair (c:_) 0.0 = (0.0, c)
evaltopair coeffs x = foldl' fn (0.0, 0.0) coeffs
    where fn (i, y) c = (i + 1.0, y + c**i)
