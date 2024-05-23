// 218. The Skyline Problem

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

// use std::{thread, time::Duration};
use std::cmp::Ordering;
use std::cmp::{max, min};
use std::collections::BinaryHeap;

#[derive(Eq, PartialEq, Debug, Clone)]
struct Build {
    left: i32,
    right: i32,
    height: i32,
}

impl Ord for Build {
    fn cmp(&self, other: &Self) -> Ordering {
        other
            .left
            .cmp(&self.left)
            .then_with(|| other.right.cmp(&self.right))
            .then_with(|| self.height.cmp(&other.height))
    }
}

impl PartialOrd for Build {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Solution {
    pub fn get_skyline(buildings: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut heap = BinaryHeap::new();
        let mut mright = 0;

        for b in &buildings {
            heap.push(Build {
                left: b[0],
                right: b[1],
                height: b[2],
            });
            mright = max(mright, b[1]);
        }
        heap.push(Build {
            left: mright,
            right: mright,
            height: 0,
        });

        let mut last = Build {
            left: -1,
            right: -1,
            height: 0,
        };

        let mut ret = vec![];
        if let Some(b) = heap.peek() {
            ret.push(vec![b.left, b.height]);
            last.left = b.left;
            last.right = b.left;
            last.height = b.height;
        }

        while let Some(b) = heap.peek() {
            let b = b.clone();
            if last.right <= b.left {
                if let Some(r) = ret.last() {
                    let (l, h) = (r[0], r[1]);
                    if l == last.left && h < last.height {
                        ret.pop();
                    }
                    if !(l <= last.left && h == last.height) {
                        ret.push(vec![last.left, last.height]);
                    }
                }

                if last.right != b.left {
                    ret.push(vec![last.right, 0]);
                }

                last = b.clone();
                heap.pop();
            } else if last.right > b.left {
                if last.height >= b.height {
                    if last.right < b.right {
                        let nb = Build {
                            left: last.right,
                            right: b.right,
                            height: b.height,
                        };
                        heap.push(nb);
                    }
                    heap.pop();
                } else {
                    if last.right > b.right {
                        let nb = Build {
                            left: b.right,
                            right: last.right,
                            height: last.height,
                        };
                        heap.push(nb);
                    }
                    last.right = b.left;
                }
            }
            // println!("at {:?}    {}  {}", last, last.left, last.height);
        }
        ret.push(vec![last.right, 0]);
        return ret;
    }
}

fn main() {
    let cases = [
        (mat![
            [2, 9, 10],
            [3, 7, 15],
            [5, 12, 12],
            [15, 20, 10],
            [19, 24, 8]
        ]),
        (mat![[0, 2, 3], [2, 5, 3]]),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::get_skyline(c.clone());
        println!("result: {:?}", r);
    }
}
