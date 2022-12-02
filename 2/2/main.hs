

main = do
    s <- readFile "input.txt"
    print  (foldr (+) 0 [ totalRoundScore (head ((words line)!!0)) (head ((words line)!!1)) | line <- lines s])


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
totalRoundScore a x 
    | (x == 'X' && a == 'A') = (roundScore a 'Z' ) + (charScore 'Z')
    | (x == 'X' && a == 'B') = (roundScore a 'X' ) + (charScore 'X')
    | (x == 'X' && a == 'C')= (roundScore a 'Y' ) + (charScore 'Y')
    | (x == 'Y' && a == 'A') = (roundScore a 'X' ) + (charScore 'X')
    | (x == 'Y' && a == 'B') = (roundScore a 'Y' ) + (charScore 'Y')
    | (x == 'Y' && a == 'C') = (roundScore a 'Z' ) + (charScore 'Z')
    | (x == 'Z' && a == 'A') = (roundScore a 'Y' ) + (charScore 'Y')
    | (x == 'Z' && a == 'B') = (roundScore a 'Z' ) + (charScore 'Z')
    | (x == 'Z' && a == 'C') = (roundScore a 'X' ) + (charScore 'X')