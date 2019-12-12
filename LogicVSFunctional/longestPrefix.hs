
longestPrefix :: [String] -> String
longestPrefix (currentPrefix:lastString:[]) = prefixMatcher currentPrefix lastString
longestPrefix (currentPrefix:nextString:remainingStrings) = longestPrefix (newPrefix:remainingStrings)
    where newPrefix = prefixMatcher currentPrefix nextString

prefixMatcher :: String -> String -> String
prefixMatcher [] _ = ""
prefixMatcher (currentPrefix:remainingPrefix) (currentChar:remainingChars)
    | currentPrefix == currentChar = [currentPrefix] ++ prefixMatcher remainingPrefix remainingChars 
    | otherwise = ""
