package leetcode.clone_graph;

import java.util.HashMap;
import java.util.LinkedList;

class Solution {
    public Node cloneGraph(Node node) {
        if (node == null) {
            return null;
        }
        Node r = null;
        HashMap<Node, Node> m = new HashMap<>();
        LinkedList<Node> willDeal = new LinkedList<>();
        willDeal.add(node);
        while (true) {
            if (willDeal.isEmpty()) {
                break;
            }
            Node old = willDeal.remove();
            Node now = new Node();
            now.val = old.val;
            now.neighbors = new LinkedList<>();
            m.put(old, now);
            if (r == null) {
                r = now;
            }

            for (Node neighbor : old.neighbors) {
                if (m.getOrDefault(neighbor, null) == null) {
                    willDeal.add(neighbor);
                }
            }
        }
        HashMap<Node, Boolean> dealed = new HashMap<>();
        willDeal = new LinkedList<>();
        willDeal.add(node);
        while (true) {
            if (willDeal.isEmpty()) {
                break;
            }
            Node old = willDeal.remove();
            Node now = m.get(old);
            if (dealed.getOrDefault(old, false)) {
                continue;
            }
            dealed.put(old, true);
            for (Node neighbor : old.neighbors) {
                if (!dealed.getOrDefault(neighbor, false)) {
                    willDeal.add(neighbor);
                }
                now.neighbors.add(m.get(neighbor));
            }
        }
        return r;
    }

    public static void main(String[] args) {
        Node t1 = new Node();
        t1.val = 1;
        Node t2 = new Node();
        t2.val = 2;
        Node t3 = new Node();
        t3.val = 3;
        Node t4 = new Node();
        t4.val = 4;

        t1.neighbors = new LinkedList<>();
        t1.neighbors.add(t2);
        t1.neighbors.add(t3);

        t2.neighbors = new LinkedList<>();
        t2.neighbors.add(t1);
        t2.neighbors.add(t4);

        t3.neighbors = new LinkedList<>();
        t3.neighbors.add(t1);
        t3.neighbors.add(t4);

        t4.neighbors = new LinkedList<>();
        t4.neighbors.add(t2);
        t4.neighbors.add(t3);

        Solution s = new Solution();
        s.cloneGraph(t1);
    }
}
