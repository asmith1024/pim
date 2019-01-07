module Poly (polY) where

import Data.List (foldl')

polY :: (Eq a, Floating a) => a -> [a] -> a
polY x cs = snd (psum x cs) 

psum :: (Eq a, Floating a) => a -> [a] -> (a, a)
psum 0.0 (c:_) = (0.0, c)
psum x coeffs = foldl' fn (0.0, 0.0) coeffs
    where fn (i, y) c = (i + 1.0, y + x**i * c)
