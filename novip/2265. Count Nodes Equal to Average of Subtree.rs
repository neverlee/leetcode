// 2265. Count Nodes Equal to Average of Subtree

struct Solution {}

#[warn(unused_macros)]
macro_rules! mat {
    ($elem:expr; $r:expr, $c:expr) => (
        vec![vec![$elem; $c]; $r]
    );
    ($($x:expr),*) => (
        [$($x),*].iter().map(|v| v.to_vec()).collect::<Vec<_>>()
    );
    ($($x:expr,)*) => (mat![$($x),*])
}

#[warn(unused_macros)]
macro_rules! vecs {
    ($($x:expr),*) => (
        vec![$($x),*].into_iter().map(|s| s.to_string()).collect::<Vec<_>>()
    );
    ($($x:expr,)*) => (vecs![$($x),*])
}

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

struct Recurrion(i32);

impl Recurrion {
    fn refn(&mut self, root: Option<Rc<RefCell<TreeNode>>>) -> (i32, i32) {
        if let Some(rcn) = root {
            let rn = rcn.borrow();
            let left = self.refn(rn.left.clone());
            let right = self.refn(rn.right.clone());
            let sum = rn.val + left.0 + right.0;
            let count = left.1 + right.1 + 1;
            if sum / count == rn.val {
                self.0 += 1;
            }
            return (sum, count);
        }
        (0, 0)
    }
}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn average_of_subtree(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut rc = Recurrion(0);
        rc.refn(root);
        rc.0
    }
}

fn main() {
    let cases = [(vec![1])];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        // let r = Solution::shortest_path_length(c.clone());
        // println!("result: {:?}", r);
    }
}
