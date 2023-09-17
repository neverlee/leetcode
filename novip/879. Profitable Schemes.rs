// 879. Profitable Schemes

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

impl Solution {
    pub fn profitable_schemes(n: i32, min_profit: i32, group: Vec<i32>, profit: Vec<i32>) -> i32 {
        let ng = group.len();
        let m = 1e9 as i32+7;

        let mut fv = vec![0; n as usize +1];
        fv[0] = 1;
        for g in group.iter() {
            for i in (0..=(n as i32-g)).rev() {
                fv[(i+g) as usize] = (fv[(i+g) as usize] + fv[i as usize])%m;
            }
        }

        let mut fm = mat![0; n as usize + 1, min_profit as usize + 1];
        fm[0][0] = 1;

        for k in 0..ng {
            let (gp, pf) = (group[k], profit[k]);
            for i in (0..=(n as i32 - gp)).rev() {
                for j in (0..min_profit-pf).rev() {
                    fm[(i+gp) as usize][(j+pf) as usize] = (fm[(i+gp) as usize][(j+pf) as usize] + fm[i as usize][j as usize])%m;
                }
            }
        }

        let (mut nall, mut nmin) = (0, 0);
        for i in 0..=n as usize {
            nall = (nall + fv[i]) % m;
        }
        for i in 0..=n as usize {
            for j in 0..min_profit as usize {
                nmin = (nmin + fm[i][j])%m;
            }
        }
        let ret = (nall+m-nmin)%m;
        println!(">>> {:?}", fv);
        
        ret 
    }
}

fn main() {
    let cases = [(
        5, 3, vec![2,2], vec![2,3],
    ), (
        5, 3, vec![2,2,0], vec![2,3,0],
    ), (
        10, 5, vec![2,3,5], vec![6,7,8],
    )];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::profitable_schemes(c.0, c.1, c.2.clone(), c.3.clone());
        println!("result: {:?}", r);
    }
}
