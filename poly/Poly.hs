module Poly (polY, printPoly) where

import Data.List (foldl')

-- polY corresponds to poly.Y in poly.go
polY :: (Eq a, Floating a) => a -> [a] -> a
polY x cs = snd (psum x cs) 

psum :: (Eq a, Floating a) => a -> [a] -> (a, a)
psum 0.0 (c:_) = (0.0, c)
psum x coeffs = foldl' fn (0.0, 0.0) coeffs
    where fn (e, y) 0 = (e + 1.0, y)
          fn (e, y) c = (e + 1.0, y + x**e * c)

-- interpolatePoly corresponds to poly.Interpolate in poly.go


-- printPoly corresponds to poly.Print in poly.go
printPoly :: [Double] -> String
printPoly []  = ""
printPoly cs  = printTerms $ getTerms cs

data Term = Term Double Int -- Coefficient, exponent

getTerms :: [Double] -> [Term]
getTerms cs = filter keepTerm (snd (termsFromCoefficients cs))

keepTerm :: Term -> Bool
keepTerm (Term c _) = c /= 0.0

termsFromCoefficients :: [Double] -> (Int, [Term])
termsFromCoefficients = foldl' nextTerm (0, [])

nextTerm :: (Int, [Term]) -> Double -> (Int, [Term])
nextTerm (e, ts) c = (e + 1, Term c e:ts)

printTerms :: [Term] -> String
printTerms []     = ""
printTerms (t:ts) = "f(x) = " ++ showFirstTerm t ++ showRemainingTerms ts

showFirstTerm :: Term -> String
showFirstTerm (Term c 0) = show c
showFirstTerm (Term c 1) = show c ++ "x"
showFirstTerm (Term c e) = show c ++ "x^" ++ show e

showRemainingTerms :: [Term] -> String
showRemainingTerms = foldl' (\s t -> s ++ showTerm t) ""

showTerm :: Term -> String
showTerm (Term c 0) = showCoefficient c
showTerm (Term c 1) = showCoefficient c ++ "x"
showTerm (Term c e) = showCoefficient c ++ "x^" ++ show e

showCoefficient :: Double -> String
showCoefficient c
    | c < 0.0   = " - " ++ show (abs c)
    | otherwise = " + " ++ show c



