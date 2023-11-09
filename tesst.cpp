
struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};
class Solution {
public:
    int minDepth(TreeNode* root) {
        // left child is null
        if (root == nullptr) return 0;
        // right child is null
        if (root->left == nullptr) return minDepth(root->right) + 1;
        // left child is null
        if (root->right == nullptr) return minDepth(root->left) + 1;
        // both children are not null
        return min(minDepth(root->left), minDepth(root->right));
    }
};