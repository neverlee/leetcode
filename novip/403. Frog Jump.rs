// 403. Frog Jump

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

#[warn(unused_macros)]
macro_rules! ifv {
    ($c:expr, $a:expr, $b: expr) => {
        if $c {
            $a
        } else {
            $b
        }
    };
}

use std::collections::{HashMap, HashSet};

impl Solution {
    pub fn can_cross(stones: Vec<i32>) -> bool {
        if stones[1] != 1 {
            return false;
        }

        let mut ha: HashMap<i32, HashSet<i32>> = HashMap::new();

        for &e in &stones {
            ha.insert(e, HashSet::new());
        }

        if let Some(aset) = ha.get_mut(&1) {
            aset.insert(1);
        }

        for i in 1..stones.len() - 1 {
            let cur = stones[i];
            if let Some(curset) = ha.remove(&cur) {
                for s in curset {
                    for si in s - 1..=s + 1 {
                        if si > 0 {
                            if let Some(aset) = ha.get_mut(&(cur + si)) {
                                aset.insert(si);
                            }
                        }
                    }
                }
            }
        }

        ha.get(stones.last().unwrap())
            .map_or(false, |s| s.len() > 0)
    }
}

fn main() {
    let cases = [
        vec![0, 1, 3, 5, 6, 8, 12, 17],
        vec![0, 1, 2, 3, 4, 8, 9, 11],
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::can_cross(c.clone());
        println!("result: {:?}", r);
    }
}
