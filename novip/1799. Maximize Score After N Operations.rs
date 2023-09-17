// 1799. Maximize Score After N Operations

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

fn gcd(a: i32, b: i32) -> i32 {
    let (mut a, mut b) = (a, b);
    let mut t;
    while b > 0 {
        t = a % b;
        a = b;
        b = t;
    }
    return a;
}

fn range_incr(sz: usize, mut f: impl FnMut(usize, usize) -> ()) {
    for i in 0..sz {
        for j in i + 1..sz {
            f(i, j);
        }
    }
}

impl Solution {
    pub fn max_score(nums: Vec<i32>) -> i32 {
        let sz = nums.len();
        let mut gcds = mat![0; sz, sz];

        range_incr(sz, |i, j| gcds[i][j] = gcd(nums[i], nums[j]));

        let fsz = 1 << sz;
        let (mut f1, mut f2) = (vec![-1; fsz], vec![-1; fsz]);
        let (mut p1, mut p2) = (&mut f1, &mut f2);
        p1[0] = 0;

        for k in 1..=(sz / 2) {
            for (ci, &n) in p1.iter().enumerate() {
                if n == -1 {
                    continue;
                }
                range_incr(sz, |x, y| {
                    let xy = (1 << x) | (1 << y);
                    if ci & xy == 0 {
                        let ni = ci | xy;
                        p2[ni] = p2[ni].max(p1[ci] + gcds[x][y] * k as i32);
                    }
                });
            }

            std::mem::swap(&mut p1, &mut p2);
        }

        p1[fsz - 1]
    }
}

fn main() {
    let cases = [(vec![1, 2]), (vec![3, 4, 6, 8]), (vec![1, 2, 3, 4, 5, 6])];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::max_score(c.clone());
        println!("result: {:?}", r);
    }
}
