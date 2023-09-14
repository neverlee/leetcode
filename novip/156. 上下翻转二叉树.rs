// 156. 上下翻转二叉树

struct Solution {}

#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

fn output_tree(deep:i32, root: &Option<Rc<RefCell<TreeNode>>>) {
    match root {
        Some(rnode) => {
            let (v, l, r) = {
                let brn = rnode.borrow();
                (brn.val, brn.left.is_some(), brn.right.is_some())
            };
            if l || r {
                print!("{}", v);
                print!("_(");
                output_tree(deep+1, &rnode.borrow().left);
                output_tree(deep+1, &rnode.borrow().right);
                print!(") ");
            } else {
                print!("{} ", v);
            }
        },
        None => print!("n "),
    }
    if deep == 0 {
        println!("");
    }
}

fn create_tree(nums:&[i32], i:usize) ->Option<Rc<RefCell<TreeNode>>>{
    if i < nums.len() {
        let num = nums[i];
        let left = create_tree(nums, i+1);
        let right = if left.is_none() { None}  else {
                Some(Rc::new(RefCell::new(TreeNode::new(nums[i]+11))))
            };

        return Some(Rc::new(RefCell::new(TreeNode{
            val: num,
            left: left,
            right: right,
        })));
    }
    None
}

use std::cell::RefCell;
use std::rc::Rc;

impl Solution {
    pub fn upside_down_binary_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        let (froot, _) = Self::upside_down_binary_tree_helper(root);
        froot
    }

    fn upside_down_binary_tree_helper(root: Option<Rc<RefCell<TreeNode>>>) -> (Option<Rc<RefCell<TreeNode>>>, Option<Rc<RefCell<TreeNode>>>) {
        if let Some(ref sroot) = &root {
            let (left, right) = {
                let mut bsroot = sroot.borrow_mut();
                (bsroot.left.take(), bsroot.right.take())
            };

            if left.is_some() {
                let (fnroot, subroot) = Self::upside_down_binary_tree_helper(left);

                if let Some(ssubroot) = subroot {
                    let mut bssubroot = ssubroot.borrow_mut();
                    bssubroot.left = right;
                    bssubroot.right = root.clone();
                };

                return (fnroot, root);
            } else {
                let last = root.clone();
                return (root, last);
            }
        };

        (None, None)
    }

}


fn main() {
    let cases = vec![3];
    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let tree = create_tree(&[1,2,3,4,5], 0);
        output_tree(0, &tree);
        let ntree = Solution::upside_down_binary_tree(tree);
        output_tree(0, &ntree);
    }
}

