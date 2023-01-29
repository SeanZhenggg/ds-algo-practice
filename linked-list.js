class ListNode {
  val;
  next;
  constructor(val, next = null) {
    this.val = val;
    this.next = next;
  }
}

function printLinkedList(list, funcName = 'linkedList') {
  let result = [];
  while (list != null) {
    result.push(list.val);
    list = list.next;
  }

  console.log(`⛔️⛔️⛔️⛔️⛔️ ${funcName} ⛔️⛔️⛔️⛔️⛔️`, result.join('->'));
}

function getLinkedListLength(list) {
  let listLen = 0;

  while (list != null) {
    listLen += 1;
    list = list.next;
  }
  return listLen;
}

function createLinkedList(data) {
  const dummyHead = new ListNode(null);
  let head = dummyHead;

  for(let i = 0; i < data.length; i++) {
    head.next = new ListNode(data[i]);
    head = head.next;
  }
  return dummyHead.next;
}

let linkedList = createLinkedList([1,3,4,8,10]);

printLinkedList(linkedList);

function addLinkedListFront(list, newVal) {
  const newNode = new ListNode(newVal);
  newNode.next = list;
  return newNode
}

linkedList = addLinkedListFront(linkedList, 2);

// console.log('⛔️⛔️⛔️⛔️⛔️ linkedList head', linkedList);

printLinkedList(linkedList, 'addLinkedListFront');

function addLinkedListNode(list, index, newVal) {
  let dummyHead = new ListNode(null);
  dummyHead.next = list;
  let curIndex = 0;
  let preHead = dummyHead; 
  while (curIndex < index) {
    preHead = list;
    list = list.next;
    curIndex += 1;
  }

  const newNode = new ListNode(newVal);
  preHead.next = newNode;
  newNode.next = list;

  return dummyHead.next;
}

linkedList = addLinkedListNode(linkedList, 2, 9);

printLinkedList(linkedList, 'addLinkedListNode');

function addLinkedListBack(list, newVal) {
  let head = list;
  const newNode = new ListNode(newVal);
  while (list.next != null) { list = list.next; }
  list.next = newNode;
  return head;
}

linkedList = addLinkedListBack(linkedList, 15);

printLinkedList(linkedList, 'addLinkedListBack');

function removeLinkedListFront(list) {
  return list.next;
}

linkedList = removeLinkedListFront(linkedList);

printLinkedList(linkedList, 'removeLinkedListFront');

function removeLinkedListNode(list, index) {
  let dummyHead = new ListNode(null);
  dummyHead.next = list;
  let curIndex = 0;
  let prevHead = dummyHead;
  while (curIndex < index) {
    prevHead = list;
    list = list.next;
    curIndex += 1;
  }

  prevHead.next = list.next;

  return dummyHead.next;
}

linkedList = removeLinkedListNode(linkedList, 2);

printLinkedList(linkedList, 'removeLinkedListNode');

function removeLinkedListBack(list) {
  let dummyHead = new ListNode(null);
  dummyHead.next = list;
  while (list.next.next != null) {
    list = list.next;
  }
  list.next = null;
  return dummyHead.next;
}

linkedList = removeLinkedListBack(linkedList)

printLinkedList(linkedList, 'removeLinkedListBack');