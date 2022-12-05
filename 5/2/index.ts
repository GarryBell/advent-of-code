import { readFileSync } from 'fs';
const file = readFileSync('./input.txt', 'utf-8');
const start = file.split('\n\n')[0]

let lines = start.split('\n')

lines = lines.map((n) => n.replace(/    /g, '   '))

lines = lines.map((n) => n.replace(/   /g, '[1]'))
lines = lines.map((n) => n.replace(/ /g, ''))

lines.pop()

const end = file.split('\n\n')[1]

const linesEnd = end.split('\n')


let array: string[][] = []
for (const row in lines) {
    let items = lines[row].split('[')
    items = items.map((n) => n.replace(/]/g, ''))
    for (const item in items) {
        if (!array[item]) {
            array[item] = []
        }
        if (items[item] !== '1') {
            array[item].push(items[item])
        }

    }

}
array.map((n) => n.reverse())
for (const row in linesEnd) {
    const split = linesEnd[row].split(' ')
    const amount = parseInt(split[1])
    const from = parseInt(split[3])
    const to = parseInt(split[5])
    let localItems = array[from].splice(-amount, amount)
    // localItems = localItems.reverse()
    for (const localItem in localItems) {
        array[to].push(localItems[localItem])
    }
}


const answer = array.map((n) => n.reverse()[0])
console.log(answer.reduce((a, b) => a + b))