// 97. Interleaving String

struct Solution {}

#[warn(unused_macros)]
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
    pub fn is_interleave(s1: String, s2: String, s3: String) -> bool {
        if s3.len() != s1.len() + s2.len() {
            return false;
        }
        let (s1, s2, s3) = (s1.as_bytes(), s2.as_bytes(), s3.as_bytes());
        let mut f = mat!(false; s1.len()+1, s2.len()+2);
        f[0][0] = true;
        for i in 0..s1.len() {
            f[i + 1][0] = f[i][0] && s1[i] == s3[i];
        }
        for i in 0..s2.len() {
            f[0][i + 1] = f[0][i] && s2[i] == s3[i];
        }
        for i in 0..s1.len() {
            for j in 0..s2.len() {
                f[i + 1][j + 1] = (f[i][j + 1] && s1[i] == s3[i + j + 1])
                    || (f[i + 1][j] && s2[j] == s3[i + j + 1]);
            }
        }
        f[s1.len()][s2.len()]
    }
}

fn main() {
    let cases = [
        ("aabcc", "dbbca", "aadbbcbcac"),
        ("aabcc", "dbbca", "aadbbbaccc"),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::is_interleave(c.0.to_string(), c.1.to_string(), c.2.to_string());
        println!("result: {:?}", r);
    }
}
