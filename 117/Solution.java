package leetcode.populating_next_right_pointers_in_each_node_ii;

public class Solution {
    public static void main(String[] args) {
        Solution s = new Solution();

        TreeLinkNode t0 = new TreeLinkNode(1);
        t0.left = new TreeLinkNode(2);
        t0.right = new TreeLinkNode(3);
        t0.right.left = new TreeLinkNode(4);
        s.connect(t0);
        s.print(t0);

        TreeLinkNode t3 = new TreeLinkNode(1);
        t3.left = new TreeLinkNode(2);
        t3.right = new TreeLinkNode(3);
        t3.left.left = new TreeLinkNode(4);
        t3.left.right = new TreeLinkNode(5);
        t3.right.right = new TreeLinkNode(6);
        t3.left.left.left = new TreeLinkNode(7);
        t3.right.right.right = new TreeLinkNode(8);
        s.connect(t3);
        s.print(t3);

        TreeLinkNode t2 = new TreeLinkNode(1);
        t2.left = new TreeLinkNode(2);
        t2.right = new TreeLinkNode(3);
        t2.left.left = new TreeLinkNode(4);
        t2.right.right = new TreeLinkNode(5);
        s.connect(t2);
        s.print(t2);

        TreeLinkNode t1 = new TreeLinkNode(1);
        t1.left = new TreeLinkNode(2);
        t1.right = new TreeLinkNode(3);
        t1.left.right = new TreeLinkNode(4);
        s.connect(t1);
        s.print(t1);
    }

    public void print(TreeLinkNode tmp) {
        TreeLinkNode ne;
        while (tmp != null) {
            TreeLinkNode tmp1 = tmp;
            ne = null;
            while (tmp1 != null) {
                System.out.print(tmp1.val + ",");
                if (ne == null) {
                    if (tmp1.left != null) {
                        ne = tmp1.left;
                    }
                }
                if (ne == null) {
                    if (tmp1.right != null) {
                        ne = tmp1.right;
                    }
                }
                tmp1 = tmp1.next;
            }
            System.out.println();
            tmp = ne;
        }
        System.out.println("+++++");
    }

    public void connect(TreeLinkNode root) {
        TreeLinkNode prev;
        TreeLinkNode node;
        TreeLinkNode nextStart = root;
        while (nextStart != null) {
            prev = null;
            node = nextStart;
            while (node != null) {
                if (node.left != null) {
                    if (prev == null) {
                        prev = node.left;
                        nextStart = node.left;
                    } else {
                        prev.next = node.left;
                        prev = node.left;
                    }
                }
                if (node.right != null) {
                    if (prev == null) {
                        prev = node.right;
                        nextStart = node.right;
                    } else {
                        prev.next = node.right;
                        prev = node.right;
                    }
                }
                node = node.next;
            }
            if (prev == null) {
                break;
            }
        }
    }
}