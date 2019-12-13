
longestPrefix :: [String] -> String

longestPrefix (currentPrefix:lastString:[]) = prefixMatcher currentPrefix lastString
longestPrefix (currentPrefix:nextString:remainingStrings) = longestPrefix (updatedPrefix:remainingStrings)
    where updatedPrefix = prefixMatcher currentPrefix nextString

prefixMatcher :: String -> String -> String

prefixMatcher [] _ = ""
prefixMatcher (currentPrefix:remainingPrefix) (currentChar:remainingChars)
    | currentPrefix == currentChar = [currentPrefix] ++ prefixMatcher remainingPrefix remainingChars 
    | otherwise = ""

