// 2251. Number of Flowers in Full Bloom

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

fn range_incr(sz: usize, mut f: impl FnMut(usize, usize) -> ()) {
    for i in 0..sz {
        for j in i + 1..sz {
            f(i, j);
        }
    }
}

use std::cmp::Reverse;
use std::collections::BinaryHeap;

impl Solution {
    pub fn full_bloom_flowers(flowers: Vec<Vec<i32>>, people: Vec<i32>) -> Vec<i32> {
        let mut pindex = (0..people.len()).collect::<Vec<_>>();
        pindex.sort_by_key(|&k| people[k]);

        let mut flowers = flowers.clone();
        flowers.sort_by(|a, b| a[0].cmp(&b[0]).then_with(|| a[1].cmp(&b[1])));

        let mut ret = people.clone();
        let mut heap = BinaryHeap::new();
        let mut cur = 0;
        for pi in pindex {
            let p = people[pi];
            while let Some(fv) = flowers.get(cur) {
                if fv[0] <= p && p <= fv[1] {
                    heap.push(Reverse(fv[1]));
                    cur += 1;
                } else if fv[0] < p {
                    cur += 1;
                    continue;
                } else {
                    break;
                }
            }

            while let Some(&k) = heap.peek() {
                if k.0 < p {
                    heap.pop();
                } else {
                    break;
                }
            }
            ret[pi] = heap.len() as i32;
        }

        ret
    }
}

fn main() {
    let cases = [
        (mat![[1, 6], [3, 7], [9, 12], [4, 13]], vec![2, 3, 7, 11]),
        (mat![[1, 10], [3, 3]], vec![3, 3, 2]),
        (
            mat![[11, 11], [24, 46], [3, 25], [44, 46]],
            vec![1, 8, 26, 7, 43, 26, 1],
        ),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::full_bloom_flowers(c.0.clone(), c.1.clone());
        println!("result: {:?}", r);
    }
}
