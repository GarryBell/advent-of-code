import { readFileSync } from 'fs';
const file = readFileSync('./input.txt', 'utf-8');
const start = file.split('\n\n')[0]

let lines = start.split('\n')


class node {
    name: string
    total: number
    children: node[]
    prevVal?: node
    constructor(name: string, total: number, children: node[], prevVal?: node) {
        this.name = name
        this.total = total
        this.prevVal = prevVal
        this.children = []

    }

}

const startNode: node = new node('/', 0, [])
let currentNode = startNode

for (const i in lines) {

    const line = lines[i]

    if (!(line.startsWith("$"))) {
        if (line.substring(0, 3) === 'dir') {
            const title = line.split(' ')[1]
            if (!currentNode.children.find((n) => n.name === title)) {
                currentNode.children.push(new node(title, 0, [], currentNode))
            }
        } else {
            const title = line.split(' ')[1]
            if (!currentNode.children.find((n) => n.name === title)) {
                currentNode.children.push(new node(title, parseInt(line.split(' ')[0]), [], currentNode))
            }
        }
    }
    if (line.startsWith("$ cd")) {
        if (line === "$ cd ..") {
            if (currentNode.prevVal) {
                currentNode = currentNode.prevVal
            }
        } else {
            const newNode = currentNode.children.find((n) => n.name === line.split(' ')[2])
            if (newNode) {
                currentNode = newNode
            }
        }
    }
}


let total = 0
let smallestVal = 70000000
let smallestName = ''

const visit = (currentNode: node) => {
    if (currentNode.children.length === 0) {
        return currentNode.total
    }
    let localTotal = currentNode.total

    for (const child in currentNode.children) {

        const res = visit(currentNode.children[child])
        if (res) {
            localTotal += res
        }
    }
    total += localTotal
    if (localTotal > 30000000 && localTotal < smallestVal) {
        smallestName = currentNode.name
        smallestVal = localTotal
    }
    return localTotal
}


const visit2 = (currentNode: node) => {
    if (currentNode.children.length === 0) {
        return currentNode.total
    }
    let localTotal = currentNode.total

    for (const child in currentNode.children) {


        const res = visit2(currentNode.children[child])
        if (res) {
            localTotal += res
        }
    }

    if (localTotal > (total - 40000000) && localTotal < smallestVal) {
        smallestName = currentNode.name

        smallestVal = localTotal
    }
    return localTotal
}


total = visit(startNode)
visit2(startNode)


console.log(smallestVal)