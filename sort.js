const _swap = (arr, i, j) => {
  let temp = arr[i]
  arr[i] = arr[j]
  arr[j] = temp
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

const input = [2,7,5,9,6,1]

console.log('input = ', input, ', bubbleSort = ', bubbleSort(input))
console.log('input = ', input, ', selectionSort = ', selectionSort(input))
console.log('input = ', input, ', insertionSort = ', insertionSort(input))