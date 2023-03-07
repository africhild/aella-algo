class Node {
    constructor(value) {
        this.value = value
        this.next = null
    }
}
const a = new Node("a")
const b = new Node("b")
const c = new Node("c")
const d = new Node("d")

a.next = b
b.next = c
c.next = d

const printNodes = (head) => {
    while(head !== null){
        console.log(head.value)
        head = head.next
    }
}
printNodes(a)