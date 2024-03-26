// 2265. Count Nodes Equal to Average of Subtree

struct Solution {}

#[allow(unused_macros)]
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

// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//     pub val: i32,
//     pub left: Option<Rc<RefCell<TreeNode>>>,
//     pub right: Option<Rc<RefCell<TreeNode>>>,
// }
// 
// impl TreeNode {
//     #[inline]
//     pub fn new(val: i32) -> Self {
//         TreeNode {
//             val,
//             left: None,
//             right: None,
//         }
//     }
// }

use std::{thread, time::Duration};
use std::cmp::{min, max};

impl Solution {
    pub fn maximum_score(nums: Vec<i32>, k: i32) -> i32 {
        let ku = k as usize;
        let (mut left, mut right) = (k-1, k+1);
        let mut cur = nums[ku];

        let mut ret = cur;
        while cur > 0 {
            while left >= 0 && nums[left as usize] >= cur {
                left -= 1;
            }
            while (right as usize) < nums.len() && nums[right as usize] >= cur {
                right += 1;
            }

            ret = max(ret, (right - left -1)*cur);

            let lv = if left >= 0 { nums[left as usize]} else {-1};
            let rv = if (right as usize) < nums.len() { nums[right as usize] } else {-1};
            cur = max(lv, rv);
        }

        ret
    }
}

fn main() {
    let cases = [(vec![1,4,3,7,4,5], 3), (vec![5,5,4,5,4,1,1,1], 0)];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::maximum_score(c.0.clone(), c.1);
        println!("result: {:?}", r);
    }
}
