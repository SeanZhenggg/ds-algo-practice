const _swap = (arr, i, j) => {
  let temp = arr[i]
  arr[i] = arr[j]
  arr[j] = temp
}

function _quickSortPartition(array, start, end) {
  const pivot = array[end];
  let i = start - 1;
  for(j = start; j <= end - 1; j ++) {
    if(array[j] <= pivot) {
      i = i + 1;
      _swap(array, i, j);
    }
  }
  _swap(array, i + 1, end);
  return i + 1;
}

function _shuffle(arr) {
  let ret = []
  let randIdx
  let set = new Set()
  for(let i = 0; i < arr.length; i++) {
    do {
      randIdx = Math.floor(Math.random() * arr.length)
    } while (set.has(randIdx))
    set.add(randIdx)
    ret[i] = arr[randIdx]
  }

  return ret
}

const bubbleSort = (arr) => {
  const retArr = arr.slice()
  const n = retArr.length
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < (n - i - 1); j++) {
      if (retArr[j] > retArr[j + 1]) {
        _swap(retArr, j, j + 1)
      }
    }
  }
  return retArr
}

const selectionSort = (arr) => {
  const retArr = arr.slice()
  const n = retArr.length
  
  for(let i = 0; i < n; i++) {
    let min = i
    for(let j = i + 1; j < n; j ++) {
      if(retArr[j] < retArr[min]) min = j
    }
    if(min != i) _swap(retArr, i, min)
  }

  return retArr
}

const insertionSort = (arr) => {
  const retArr = arr.slice()
  const n = retArr.length

  for(let i = 1; i < n; i++) {
    let item = retArr[i], j = i - 1
    while(j >= 0 && retArr[j] > item) {
      retArr[j + 1] = retArr[j]
      j--;
    }

    retArr[j + 1] = item
  }

  return retArr
}

const quickSort = (array, start, end) => {
  if(start < end) {
    // partition
    var pivot = _quickSortPartition(array, start, end);
    // left
    quickSort(array, start, pivot - 1);
    // right
    quickSort(array, pivot + 1, end);
  }
}

const length = 20
const input = _shuffle(Array.from({ length }).map((_, i) => Math.floor(Math.random() * 500)))

console.log(`input \t\t= [${input}]`)
console.log(`bubbleSort \t= [${bubbleSort(input).toString()}]`)
console.log(`selectionSort \t= [${selectionSort(input).toString()}]`)
console.log(`insertionSort \t= [${insertionSort(input).toString()}]`)
console.log(`quickSort \t= [${insertionSort(input).toString()}]`)