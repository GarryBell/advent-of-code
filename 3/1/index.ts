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
for (const index in lines) {
    let lookupArray: string[] = []
    const line = lines[index]
    const first = line.substring(0, (line.length / 2))
    const second = line.substring((line.length / 2), line.length)

    for (const i in first.split('')) {
        if (second.includes(first[i]) && !(lookupArray.includes(first[i]))) {
            total += getCharVal(first[i])
            lookupArray.push(first[i])
        }
    }
}

console.log(total)