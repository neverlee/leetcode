// 1615. Maximal Network Rank

struct Solution {}

macro_rules! mat {
    ($elem:expr; $r:expr, $c:expr) => (
        vec![vec![$elem; $c]; $r]
    );
    ($($x:expr),*) => (
        // <[_]>::into_vec(box [$($x),*])
        [$($x),*].iter().map(|v| v.to_vec()).collect::<Vec<_>>()
    );
    ($($x:expr,)*) => (mat![$($x),*])
}

impl Solution {
    pub fn maximal_network_rank(n: i32, roads: Vec<Vec<i32>>) -> i32 {
        let n = n as usize;
        let mut g = mat![false; n, n];
        let mut count = vec![0; n];

        for r in roads {
            let (x, y) = (r[0] as usize, r[1] as usize);
            g[x][y] = true;
            g[y][x] = true;
            count[x] += 1;
            count[y] += 1;
        }

        let mut midx = (0..n as usize).collect::<Vec<usize>>();
        midx.sort_by_key(|ix| -count[*ix]);

        // println!("{:?}", g); println!("{:?}", count); println!("{:?}", midx);
        let mut ret = 0;
        for i in 0..n - 1 {
            for j in i + 1..n {
                let (pi, pj) = (midx[i], midx[j]);
                let r = count[pi] + count[pj] - if g[pi][pj] { 1 } else { 0 };
                if r >= ret {
                    ret = r;
                } else {
                    break;
                }
            }
        }
        ret
    }
}

fn main() {
    let cases = [
        (4, mat![[0, 1], [0, 3], [1, 2], [1, 3]]),
        (5, mat![[0, 1], [0, 3], [1, 2], [1, 3], [2, 3], [2, 4]]),
        (8, mat![[0, 1], [1, 2], [2, 3], [2, 4], [5, 6], [5, 7]]),
        (
            6,
            mat![[0, 1], [0, 2], [0, 5], [1, 4], [1, 2], [2, 5], [3, 5]],
        ),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::maximal_network_rank(c.0, c.1.clone());
        println!("result: {:?}", r);
    }
}
