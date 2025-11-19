Singly-Linked List
For this assignment, you will be implementing a Singly-Linked List with a head and tail reference. Recall that a Linked List is a collection of nodes, with each node containing a data item and reference(s) to other node(s). In a Singly-Linked List, each node will only reference the next node. Your Singly-Linked List should follow the requirements stated in the javadocs of each method you are to implement.

IMPORTANT:

You will be given 5 attempts on this assignment, with a 30 minute cooldown between submissions.

Please run your code before each submission to ensure that there are no formatting errors! If there are formatting errors in your code, your code will not be graded and a submission attempt will be logged. For more information, please review the Vocareum overview below.

Nodes

A SinglyLinkedListNode class is provided to you and will be used to represent the nodes in your Singly-Linked List. This file should be treated as read-only and should not be modified in any way. This SinglyLinkedListNode class contains getter and setter methods to access and mutate the structure of the nodes. Please make sure that you understand how this class works, as interaction with this class is crucial for this assignment.

Adding

You will implement two add() methods. One will add to the front and one will add to the back. Be sure to correctly reassign the head and/or tail pointers when adding an element.