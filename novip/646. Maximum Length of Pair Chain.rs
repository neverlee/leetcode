// 646. Maximum Length of Pair Chain

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

use std::cmp::Ordering;

impl Solution {
    pub fn find_longest_chain(pairs: Vec<Vec<i32>>) -> i32 {
        let n = pairs.len();
        let mut pairs: Vec<_> = pairs.iter().map(|e| (e[0], e[1])).collect();
        pairs.sort_by(|a, b| match a.1.cmp(&b.1) {
            Ordering::Equal => a.0.cmp(&b.0),
            or => or,
        });

        let mut f = vec![0; n];
        f[0] = 1;
        for x in 1..n {
            f[x] = f[x - 1];
            let (mut i, mut j, mut m) = (0, x as i32 - 1, 0);
            while i < j {
                m = (i + j + 1) / 2;
                if pairs[m as usize].1 >= pairs[x as usize].0 {
                    j = m - 1;
                } else {
                    i = m;
                }
            }
            let i = i as usize;
            f[x] = f[x - 1];
            if i < x && pairs[i].1 < pairs[x].0 {
                f[x] = f[x].max(f[i as usize] + 1);
            }
        }

        f[n - 1]
    }
}

fn main() {
    let cases = [
        (mat![[1, 2], [2, 3], [3, 4]]),
        (mat![[1, 2], [7, 8], [3, 4]]),
        (mat![[1, 2], [7, 8], [3, 4], [3, 5], [4, 5]]),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::find_longest_chain(c.clone());
        println!("result: {:?}", r);
    }
}
