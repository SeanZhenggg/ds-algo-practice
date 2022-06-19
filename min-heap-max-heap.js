class HeapNode {
  constructor(element, key) {
      this.element = element
      this.key = key
  }
}

class Heap {
  _heap;
  constructor() {
    this._heap = []
  }
  swap(a, b) { [this._heap[a], this._heap[b]] = [this._heap[b], this._heap[a]] }
  top = () => this._heap[0].element
  isHeapEmpty = () => this.size() < 1
  size = () => this._heap.length
  getHeap = () => this._heap
  getParentNode = node => Math.floor((node - 1) / 2)
  findPosition = element => this._heap.findIndex(node => node.element === element)
  heapify(node, length) {}
  build(array) {}
  change(element, newKey) {}
  extract() {}
}

class MinHeap extends Heap {
  constructor() {
    super()
  }
  heapify(node, length) {
    let left = (2 * node) + 1, right = (2 * node) + 2, smallest = node

    if(left <= length && this._heap[left].key < this._heap[node].key) {
      smallest = left
    }
    if(right <= length && this._heap[right].key < this._heap[smallest].key) {
      smallest = right
    }

    if(smallest !== node) {
      this.swap(smallest, node)
      this.heapify(smallest, length)
    }
  }
  build(array) {
    for(let i = 0; i < array.length; i++) {
      this._heap.push(new HeapNode(i, array[i]))
    }
    for(let i = Math.floor((this.size() - 1) / 2); i >= 0; i --) {
      this.heapify(i, this.size() - 1)
    }
  }
  change(element, newKey) {
    if(this.isHeapEmpty()) {
      console.log('heap is empty')
      return null
    }
    let index_node = this.findPosition(element)
    if(this._heap[index_node].key < newKey) {
      console.log('new key is smaller than current key')
      return;
    }

    this._heap[index_node].key = newKey

    while (index_node > 0 && this._heap[this.getParentNode(index_node)].key >= this._heap[index_node].key) {
      this.swap(index_node, this.getParentNode(index_node))
      index_node = this.getParentNode(index_node)
    }
  }
  insert(element, key) {
    this._heap.push(new HeapNode(element, key))

    this.change(element, key)
  }
  extract() {
    if(this.isHeapEmpty()) {
      console.log('heap is empty')
      return null
    }
    let min = this.top()
    this._heap[0] = this._heap[this.size() - 1]
    this._heap.pop()
    this.heapify(0, this.size() - 1)
    return min
  }
}

class MaxHeap extends Heap {
  constructor() {
    super()
  }
  heapify(node, length) {
    let left = (2 * node) + 1, right = (2 * node) + 2, biggest = node
    if(left <= length && this._heap[left].key > this._heap[node].key) {
      biggest = left
    }
    if(right <= length && this._heap[right].key > this._heap[biggest].key) {
      biggest = right
    }

    if(biggest !== node) {
      this.swap(biggest, node)
      this.heapify(biggest, length)
    }
  }
  build(array) {
    for(let i = 0; i < array.length; i++) {
      this._heap.push(new HeapNode(i, array[i]))
    }
    for(let i = Math.floor((this.size() - 1) / 2); i >= 0; i--) {
      this.heapify(i, this.size() - 1)
    }
  }
  change(element, newKey) {
    if(this.isHeapEmpty()) {
      console.log('heap is empty')
      return null
    }
    let index_node = this.findPosition(element)
    if(this._heap[index_node].key > newKey) {
      console.log('new key is larger than current key')
      return;
    }
    this._heap[index_node].key = newKey

    while(index_node > 0 && this._heap[this.getParentNode(index_node)].key < this._heap[index_node].key) {
      this.swap(index_node, this.getParentNode(index_node))
      index_node = this.getParentNode(index_node)
    }
  }
  insert(element, key) {
    this._heap.push(new HeapNode(element, key))
    this.change(element, key)
  }
  extract() {
    if(this.isHeapEmpty()) {
      console.log('heap is empty')
      return null
    }
    let max = this.top()
    this._heap[0] = this._heap[this.size() - 1]
    this._heap.pop()
    this.heapify(0, this.size() - 1) 
    return max
  }
}

let minHeap = new MinHeap()

minHeap.build([2, 8, 10, 4, 7, 20, 22, 13, 3])
console.log('heap = ', minHeap.getHeap())
minHeap.insert(9, 1)
console.log('heap after insert = ', minHeap.getHeap())
minHeap.change(5, 6)
console.log('heap after decrease = ', minHeap.getHeap())
minHeap.extract()
console.log('heap after extract min = ', minHeap.getHeap())

let maxHeap = new MaxHeap()

maxHeap.build([2, 8, 10, 4, 7, 20, 22, 13, 3])
console.log('heap = ', maxHeap.getHeap())
maxHeap.insert(9, 19)
console.log('heap after insert = ', maxHeap.getHeap())
maxHeap.change(0, 24)
console.log('heap after decrease = ', maxHeap.getHeap())
maxHeap.extract()
console.log('heap after extract min = ', maxHeap.getHeap())