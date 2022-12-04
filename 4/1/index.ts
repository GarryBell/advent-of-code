import { readFileSync } from 'fs';
const file = readFileSync('./input.txt', 'utf-8');

const lines = file.split('\n')
let total = 0

const checkRange = (first: string[], second: string[]) => {
    if (parseInt(first[0]) <= parseInt(second[0]) && parseInt(first[1]) >= parseInt(second[1])) {
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