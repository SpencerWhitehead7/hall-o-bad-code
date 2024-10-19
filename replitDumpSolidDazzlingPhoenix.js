// SOLID. DAZZLING. PHOENIX.
// what's there to say about this one? When I copied it out, there didn't happen to be much in it
// but for many years this was the go-to scratch file for anything and everything I wanted to try

function TreeNode(val, left, right) {
  this.val = val === undefined ? 0 : val;
  this.left = left === undefined ? null : left;
  this.right = right === undefined ? null : right;
}

const buildTree = (preorder, inorder) => {
  console.log(preorder, inorder);
  if (preorder.length === 0 && inorder.length === 0) {
    return null;
  }
  if (preorder.length === 1 && inorder.length === 1) {
    return new TreeNode(preorder[0]);
  }

  const rootVal = preorder[0];
  const inorderSplitPoint = inorder.indexOf(rootVal);
  const inorderL = inorder.slice(0, inorderSplitPoint);
  const inorderR = inorder.slice(inorderSplitPoint + 1);
  const lVals = new Set(inorderL);
  const rVals = new Set(inorderR);
  const preorderL = preorder.filter((v) => lVals.has(v));
  const preorderR = preorder.filter((v) => rVals.has(v));
  console.log({ inorderL, inorderR, preorderL, preorderR });

  return new TreeNode(
    preorder[0],
    buildTree(preorderL, inorderL),
    buildTree(preorderR, inorderR)
  );
};

console.log(buildTree([1, 2], [2, 1]));
