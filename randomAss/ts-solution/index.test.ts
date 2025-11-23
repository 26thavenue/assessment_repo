
import { SinglyLinkedList } from "./index.ts";

describe("SinglyLinkedList", () => {
  let list: SinglyLinkedList<number>;

  beforeEach(() => {
    list = new SinglyLinkedList<number>();
  });
  it("should add a node to the front when list is empty", () => {
    list.addFront(10);

    expect(list.getHead()?.getData()).toBe(10);
    expect(list.getTail()?.getData()).toBe(10);
    expect(list.getSize()).toBe(1);
  });

  it("should add a node to the front when list is not empty", () => {
    list.addFront(10);
    list.addFront(20);

    expect(list.getHead()?.getData()).toBe(20);
    expect(list.getTail()?.getData()).toBe(10);
    expect(list.getSize()).toBe(2);

    const headNext = list.getHead()?.getNext();
    expect(headNext?.getData()).toBe(10);
  });

  it("should add a node to the back when list is empty", () => {
    list.addBack(5);

    expect(list.getHead()?.getData()).toBe(5);
    expect(list.getTail()?.getData()).toBe(5);
    expect(list.getSize()).toBe(1);
  });

  it("should add a node to the back when list is not empty", () => {
    list.addBack(5);
    list.addBack(15);

    expect(list.getHead()?.getData()).toBe(5);
    expect(list.getTail()?.getData()).toBe(15);
    expect(list.getSize()).toBe(2);

    const next = list.getHead()?.getNext();
    expect(next?.getData()).toBe(15);
  });

  it("should maintain correct structure after mixed operations", () => {
    list.addFront(3);
    list.addBack(5);
    list.addFront(1);

    expect(list.getSize()).toBe(3);

    const head = list.getHead();
    expect(head?.getData()).toBe(1);

    const mid = head?.getNext();
    expect(mid?.getData()).toBe(3);

    const tail = mid?.getNext();
    expect(tail?.getData()).toBe(5);

    expect(list.getTail()?.getData()).toBe(5);
  });

  it("should correctly update size after many insertions", () => {
    for (let i = 0; i < 10; i++) {
      list.addBack(i);
    }

    expect(list.getSize()).toBe(10);
    expect(list.getHead()?.getData()).toBe(0);
    expect(list.getTail()?.getData()).toBe(9);
  });

});
