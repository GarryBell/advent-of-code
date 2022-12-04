import { readFileSync } from 'fs';
const file = readFileSync('./input.txt', 'utf-8');

const lines = file.split('\n')
let total = 0


// sadly I've forgotten the neater arithmetic way of doing this...
const checkRange = (first: string[], second: string[]) => {
    const first1 = parseInt(first[0])
    const first2 = parseInt(first[1])
    const second1 = parseInt(second[0])
    const second2 = parseInt(second[1])
    if (second1 <= first1 && first1 <= second2) {
        return true
    }
    if (second1 <= first2 && first2 <= second2) {
        return true
    }
}

for (const index in lines) {
    const line = lines[index]
    const words = line.split(',')

    const firstSplit = words[0].split('-')
    const secondSplit = words[1].split('-')
    if (checkRange(firstSplit, secondSplit) || checkRange(secondSplit, firstSplit)) {
        total += 1
    }


}

console.log(total)