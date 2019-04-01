package leetcode.copy_list_with_random_pointer;

import java.util.HashMap;

class Solution {
    public Node copyRandomList(Node head) {
        Node newHead = head;
        Node fake = new Node();
        Node prev = fake;
        HashMap<Node, Node> hash = new HashMap<>();
        while (true) {
            if (newHead == null) {
                break;
            }
            Node item = new Node();
            item.val = newHead.val;
            prev.next = item;
            prev = item;
            hash.put(newHead, item);

            newHead = newHead.next;
        }
        newHead = head;
        prev = fake.next;
        while (true) {
            if (newHead == null) {
                break;
            }
            if (newHead.random != null) {
                prev.random = hash.get(newHead.random);
            }

            newHead = newHead.next;
            prev = prev.next;
        }
        return fake.next;
    }
}
