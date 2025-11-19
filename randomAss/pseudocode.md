Single linked list 

means all the nodes are linked with two other nodes exceot they are head or tail nodes 
so you have node, node.next, data
if there is no node. next it is the tail and if there is no node.prev it is the head 


nodes 
^^^^^^^^^
Single Linked Node class , getter snd setter class 
get data , get next
set the next node , set the data in the node 


Add Methods 
^^^^^^^^^^^^^^^^^^^^^
add to the front 
so for this you create a new node, set the new node.next to the head  and make the new node the head 

add to the back 
Since we don't know the last node, we have to travese till we get ther so the same thing , create a new node or intitialize a new node,check if the the head is null that means there are no nodes , then oop though the list by use node.next is not null unitl you finally get to null the last node, then add a next to be the new node 
