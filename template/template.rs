//

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

struct Solution {}

impl Solution {
    pub fn add(a:i32, b:i32) -> i32 {


    }
}

fn main() {
    let cases = [
        (1,2),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::add(c.0, c.1);
        println!("result: {:?}", r);
    }
}
