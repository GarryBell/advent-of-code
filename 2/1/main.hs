

main = do
    s <- readFile "input.txt"

    print  (foldr (+) 0 [ totalRoundScore (head ((words line)!!0)) (head ((words line)!!1)) | line <- lines s])
    -- print  (totalRoundScore  'A' 'Y')

charScore::Char->Int
charScore 'X' = 1
charScore 'Y' = 2
charScore 'Z' = 3

roundScore::Char->Char->Int
roundScore  'A' 'X' = 3
roundScore  'B' 'Y' = 3
roundScore  'C' 'Z' = 3
roundScore  'A' 'Z' = 0
roundScore  'B' 'X' = 0
roundScore  'C' 'Y' = 0
roundScore  'A' 'Y' = 6
roundScore  'B' 'Z' = 6
roundScore  'C' 'X' = 6

totalRoundScore::Char->Char->Int
totalRoundScore a x = (roundScore a x ) + (charScore x)