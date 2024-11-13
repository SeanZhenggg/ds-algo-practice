function quickSelect(arr, k, type = 'largest') {
  let left = 0, right = arr.length - 1, pivot
  let kLargest = type === 'largest' ? arr.length - k : k - 1
  while(left <= right) {
    pivot = partition(arr, left, right)

    if (kLargest === pivot) return arr[pivot]

    if (kLargest > pivot) left = pivot + 1
    else right = pivot - 1
  }
}

function partition(arr, l, r) {
  let pivot = r;
  let p = l;
  for (let i = l; i < r; i++) {
    if(arr[i] < arr[pivot]) {
      swap(arr, i, p);
      p++;
    }
  }
  swap(arr, p, pivot);
  return p;
}

function swap(arr, i, j) {
  [arr[i], arr[j]] = [arr[j], arr[i]]
}

function shuffle(arr) {
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

const length = 10
const k = 2
const inputArr = shuffle(Array.from({ length }).map((_, i) => Math.floor(Math.random() * 500)))

const largest = quickSelect(inputArr.slice(), k)
console.log('the ' + `${k} largest value in [${inputArr.toString()}] = ` + largest)

const smallest = quickSelect(inputArr.slice(), k, 'smallest')
console.log('the ' + `${k} smallest value in [${inputArr.toString()}] = ` + smallest)