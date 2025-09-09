countCharFilter :: Char -> String -> Int
countCharFilter targetChar s = length (filter (== targetChar) s)

main :: IO ()
main = do
    content <- readFile "./input.txt"
    let openCount = countCharFilter '(' content
    let closeCount = countCharFilter ')' content
    let total = openCount - closeCount 
    putStrLn ( show total)
