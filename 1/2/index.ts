import { readFileSync } from 'fs';

const file = readFileSync('./input.txt', 'utf-8');

let first = 0
let second = 0
let third = 0
const handleRankingChange = (latest: number) => {
    let sortingArray = [first, second, third, latest]
    sortingArray.sort((a, b) => b - a)
    first = sortingArray[0]
    second = sortingArray[1]
    third = sortingArray[2]
}

var separatedElves = file.split('\n\n')
for (const elf in separatedElves) {
    let localMax = 0
    var items = separatedElves[elf].split('\n')
    for (const item in items) {
        if (!isNaN(parseInt(items[item]))) {
            localMax += parseInt(items[item])
        }
    }
    if (localMax > third) {
        handleRankingChange(localMax)
    }
    localMax = 0
}
console.log(first + second + third)
