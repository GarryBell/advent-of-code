import { readFileSync } from 'fs';

const file = readFileSync('input.txt', 'utf-8');

let max = 0

var separatedElves = file.split('\n\n')
for (const elf in separatedElves) {
    let localMax = 0
    var items = separatedElves[elf].split('\n')
    for (const item in items) {
        if (!isNaN(parseInt(items[item]))) {
            localMax += parseInt(items[item])
        }
    }
    if (localMax > max) {
        max = localMax
    }
    localMax = 0
}
console.log(max)
