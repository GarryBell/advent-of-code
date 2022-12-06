
main = do
    s <- readFile "input.txt"
    print((func s 0) + 4)     
    
func::[Char]->Int->Int
func (a:b:c:d:xs) index
    | (not (a == b)) && (not (b == c)) && (not (c == d)) && (not (a == c)) && ( not (a == d)) && ( not (b == d))  = index
    | otherwise = func (b:c:d:xs) (index+1)
