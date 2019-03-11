module Sets (allocatePreferences, areAllocationsStable) where

import Data.List (foldl', sortBy)

-- Our model is a list of lists of Int, where each element of a child list 
-- is an ordinal pointing at an element of the other set in the preferences
-- allocation problem.
-- 
-- While sets are unordered, we are using the list as a convenient way to
-- idenfify corresponding elements in the other set.
--
-- A preference allocation is a pair of ordinals, one from each set.
--
-- A list of preference allocations is the ultimate result of the 
-- preference allocation process. We expect this result to contain an
-- entry that uniquely combines every element from both source sets.

-- 1.0 TDD skeleton

--allocatePreferences :: [[Int]] -> [[Int]] -> [(Int, Int)]
--allocatePreferences _ _ = []

--areAllocationsStable :: [[Int]] -> [[Int]] -> [(Int, Int)] -> Bool
--areAllocationsStable _ _ _ = False

-- 1.1 First pass

allocatePreferences :: [[Int]] -> [[Int]] -> [(Int, Int)]
allocatePreferences _ _ = []

-- returns a pair, the first element of identifes a member of the other set, the second its corresponding preference ordinal 
indexPreferences :: [Int] -> [(Int, Int)]
indexPreferences ps = zipWith (\x y -> (x,y)) ps [1..]

indexEverything :: [[Int]] -> [(Int, [(Int, Int)])]
indexEverything = zipWith(\x y -> (x, indexPreferences y)) [1..]

mutualPreferences :: [[Int]] -> [[Int]] -> [((Int, Int), Int)]
mutualPreferences as bs = preferenceProduct (indexEverything as) (indexEverything bs)

-- gets the cartesion product of all possible pairings
preferenceProduct :: [(Int, [(Int, Int)])] -> [(Int, [(Int, Int)])] -> [((Int, Int), Int)]
preferenceProduct as bs = [((fst x, fst y), findProduct x y) | x <- as, y <- bs]

findProduct :: (Int, [(Int, Int)]) -> (Int, [(Int, Int)]) -> Int
findProduct as bs = preference (fst bs) (snd as) * preference (fst as) (snd bs)

preference :: Int -> [(Int, Int)] -> Int
preference i = snd . head . filter (\v -> fst v == i)

areAllocationsStable :: [[Int]] -> [[Int]] -> [(Int, Int)] -> Bool
areAllocationsStable _ _ _ = False

-- need to first calculate all possible pairings
-- then calculate the score for each
-- then return the lowest

-- list of pairs
-- GLOBAL list of preferences indexed by pair
