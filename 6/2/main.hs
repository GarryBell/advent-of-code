
main = do
    s <- readFile "input.txt"
    print((func s 0) + 14)

checkHas::[Char]->Bool
checkHas (x:xs) = (elem x xs ) || (checkHas xs)
checkHas [] = False

func::[Char]->Int->Int
func xs index
    | not (checkHas (take 14 xs))  = index
    | otherwise = func (tail xs) (index+1)
