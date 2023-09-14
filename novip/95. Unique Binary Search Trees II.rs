// 95. Unique Binary Search Trees II

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
            print!("{}", v);
            if l || r {
            print!("_(");
                output_tree(deep+1, &rnode.borrow().left);
                output_tree(deep+1, &rnode.borrow().right);
            print!(") ");
            }
        },
        None => print!("n "),
    }
    if deep == 0 {
        println!("");
    }
}

use std::collections::HashMap;
struct Helper{
    cache : HashMap<(i32, i32), Vec<Option<Rc<RefCell<TreeNode>>>>>,
}

impl Helper{
    pub fn generate_trees(&mut self, start: i32, end: i32) -> Vec<Option<Rc<RefCell<TreeNode>>>> {
        let mut res = Vec::new();
        if start > end {
            res.push(None);
            return res;
        }

        if let Some(&ref r) = self.cache.get(&(start, end)) {
            return r.to_vec();
        }

        for mid in start..(end+1) {
            let lstrees = self.generate_trees(start, mid-1);
            let rstrees = self.generate_trees(mid+1, end);
            for lt in lstrees.iter() {
                for rt in rstrees.iter() {
                    let node = TreeNode{
                        val:mid,
                        left: lt.clone(),
                        right: rt.clone(),
                    };
                    res.push(Some(Rc::new(RefCell::new(node))));
                }
            }
        }
        self.cache.insert((start, end), res.clone());

        res
    }
}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn generate_trees(n: i32) -> Vec<Option<Rc<RefCell<TreeNode>>>> {
        let mut h = Helper{ cache : HashMap::new(), };
        let ret = h.generate_trees(1, n);

        ret
    }

}


fn main() {
    let cases = vec![3];
    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        Solution::generate_trees(*c);
    }
}
