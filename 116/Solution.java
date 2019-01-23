package leetcode.populating_next_right_pointers_in_each_node;

public class Solution {
    public static void main(String[] args) {
        TreeLinkNode t = new TreeLinkNode(1);
        t.left = new TreeLinkNode(2);
        t.right = new TreeLinkNode(3);
        t.left.left = new TreeLinkNode(4);
        t.left.right = new TreeLinkNode(5);
        t.right.left = new TreeLinkNode(6);
        t.right.right = new TreeLinkNode(7);
        t.left.left.left = new TreeLinkNode(8);
        t.left.left.right = new TreeLinkNode(9);
        t.left.right.left = new TreeLinkNode(10);
        t.left.right.right = new TreeLinkNode(11);
        t.right.left.left = new TreeLinkNode(12);
        t.right.left.right = new TreeLinkNode(13);
        t.right.right.left = new TreeLinkNode(14);
        t.right.right.right = new TreeLinkNode(15);

        Solution s = new Solution();
        s.connect(t);

        TreeLinkNode tmp = t;
        while (tmp != null) {
            TreeLinkNode tmp1 = tmp;
            while (tmp1 != null) {
                System.out.println(tmp1.val);
                tmp1 = tmp1.next;
            }
            System.out.println("______");
            tmp = tmp.left;
        }
    }

    public void connect(TreeLinkNode root) {
        if (root != null && root.left != null) {
            _connect(root.left, root.right);
        }
    }

    public void _connect(TreeLinkNode node, TreeLinkNode next) {
        node.next = next;
        if (node.left != null) {
            _connect(node.left, node.right);
            _connect(node.right, next.left);
            _connect(next.left, next.right);
        }
    }
}