import { readFileSync } from 'fs';
const file = readFileSync('./input.txt', 'utf-8');

const lines = file.split('\n')
let total = 0

const getCharVal = (char: string) => {
    const val = char.charCodeAt(0)
    if (val > 96) {
        return val - 96
    } else {
        return val - 38
    }

}

const checkMembership = (a: string, b: string, c: string) => {
    for (const i in a.split('')) {
        if (b.includes(a[i]) && c.includes(a[i])) {
            return a[i]
        }
    }
}

for (let index = 0; index < lines.length; index += 3) {
    const first = lines[index]
    const second = lines[index + 1]
    const third = lines[index + 2]
    const a = checkMembership(first, second, third)
    if (a) {
        total += getCharVal(a)
    }
}

console.log(total)