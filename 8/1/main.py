

with open('input.txt') as f:
    lines = f.readlines()    
    height = len(lines)
    width = len(lines[0])
    array = [[[lines[j][i], i, j, False] for i in range(width-1)] for j in range(height)] 

def find(array2):
    for i in range(len(array2)):
        localString = array2[i]
        localMax = -1         

        for char in localString:
            if(int(char[0]) > localMax):
                localMax = int(char[0])
                char[3] = True
    
    return array2
            
def spin(array2):
    new = list(zip(*array2[::-1]))
    array2 = new
    return array2

max = 0
array = find(array)
array = spin(array)
array = find(array)
array = spin(array)
array = find(array)
array = spin(array)
array = find(array)

for i in range(len(array)):
    localString = array[i]
    localMax = 0 
    for char in localString:
        if(char[3] == True):
            max = max + 1
print(max)