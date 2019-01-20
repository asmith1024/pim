module Poly (interpolatePoly, polY, printPoly) where

import Data.List (foldl', sortBy)

-- polY corresponds to poly.Y in poly.go
polY :: Double -> [Double] -> Double
polY x = sum . zipWith (\e c -> x^^e * c) [0..]

-- interpolatePoly corresponds to poly.Interpolate in poly.go
interpolatePoly :: [(Double, Double)] -> [Double]
interpolatePoly = interpolate . validPoints

interpolate :: [(Double, Double)] -> [Double]
interpolate ps = map (\(Term c e) -> c) $ simplifyTerms (concatMap (reduceTerms . termsAtPoint ps) ps)

reduceTerms :: [[Double]] -> [Term]
reduceTerms ps = let (first:next) = map getTerms ps
                  in simplifyTerms (foldl' (\p t -> [a `mult` b | a <- p, b <- t]) first next)

termsAtPoint :: [(Double, Double)] -> (Double, Double) -> [[Double]]
termsAtPoint ps (x,y) = includeDenominator y $ processPoint x (filter (\(a,b) -> x /= a) ps)

processPoint :: Double -> [(Double, Double)] -> ([[Double]], Double)
processPoint x ps = (numerator ps, denominator x ps)

numerator :: [(Double, Double)] -> [[Double]]
numerator ps = [[(-1)*fst p, 1] | p <- ps]

denominator :: Double -> [(Double, Double)] -> Double
denominator x ps = product [x - fst p | p <- ps]

includeDenominator :: Double -> ([[Double]], Double) -> [[Double]]
includeDenominator y (ts, d) = [y/d]:ts

sumReduce :: [[Double]] -> [Double]
sumReduce = foldl' (longZip (+)) []

longZip :: (a -> a -> a) -> [a] -> [a] -> [a]
longZip _ [] bs = bs
longZip _ as [] = as
longZip fn (a:as) (b:bs) = a `fn` b : longZip fn as bs

data Term = Term Double Int

getTerms :: [Double] -> [Term]
getTerms cs = zipWith Term cs [0..]

cproduct :: [Term] -> [Term] -> [Term]
cproduct xs ys = [mult x y | x <- xs, y <- ys]

mult :: Term -> Term -> Term
mult (Term a b) (Term c d) = Term (a*c) (b+d)

simplifyTerms :: [Term] -> [Term]
simplifyTerms [] = []
simplifyTerms ts = [foldl' addTerms x xs | (x:xs) <- groupByDegree $ sortByDegree ts]

addTerms :: Term -> Term -> Term
addTerms (Term a b) (Term c _) = Term (a+c) b

groupByDegree :: [Term] -> [[Term]]
groupByDegree = foldr addTerm []

sortByDegree :: [Term] -> [Term]
sortByDegree = sortBy (\(Term _ a) (Term _ b) -> compare a b)

addTerm :: Term -> [[Term]] -> [[Term]]
addTerm t [] = [[t]]
addTerm t@(Term _ e) ts@(x@(Term _ e':xs'):xs) 
    | e == e' = (t:x):xs
    | otherwise = [t]:ts

validPoints :: [(Double, Double)] -> [(Double, Double)]
validPoints = rejectDup . sortBy (\(a, _) (b, _) -> compare a b)

rejectDup :: [(Double, Double)] -> [(Double, Double)]
rejectDup ps 
    | sortedHasDup ps = []
    | otherwise = ps

sortedHasDup :: [(Double, Double)] -> Bool
sortedHasDup = snd . foldl' (\(a, b) (c, _) -> (c, b || a == c)) (0.0, False)

-- printPoly corresponds to poly.Print in poly.go
printPoly :: [Double] -> String
printPoly = printTerms . filter notZero . getTerms

notZero :: Term -> Bool
notZero (Term c _) = c /= 0.0

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