
 export class SinglyLinkedListNode<T> {
  private data: T;
  private next: SinglyLinkedListNode<T> | null;

  constructor(data: T) {
    this.data = data;
    this.next = null;
  }

  getData(): T {
    return this.data;
  }

  getNext(): SinglyLinkedListNode<T> | null {
    return this.next;
  }

  setNext(node: SinglyLinkedListNode<T> | null): void {
    this.next = node;
  }
}


export class SinglyLinkedList<T> {
  private head: SinglyLinkedListNode<T> | null = null;
  private tail: SinglyLinkedListNode<T> | null = null;
  private size: number = 0;

  addFront(data: T): void {
    const newNode = new SinglyLinkedListNode(data);

    if (!this.head) {
      this.head = newNode;
      this.tail = newNode;
    } else {
      newNode.setNext(this.head);
      this.head = newNode;
    }

    this.size++;
  }

  addBack(data: T): void {
    const newNode = new SinglyLinkedListNode(data);

    if (!this.tail) {
      this.head = newNode;
      this.tail = newNode;
    } else {
      this.tail.setNext(newNode);
      this.tail = newNode;
    }

    this.size++;
  }

  getHead(): SinglyLinkedListNode<T> | null {
    return this.head;
  }

  getTail(): SinglyLinkedListNode<T> | null {
    return this.tail;
  }

  getSize(): number {
    return this.size;
  }
}
