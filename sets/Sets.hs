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

areAllocationsStable :: [[Int]] -> [[Int]] -> [(Int, Int)] -> Bool
areAllocationsStable _ _ _ = False

