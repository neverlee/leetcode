// 767. Reorganize String

struct Solution {}

macro_rules! mat {
    ($elem:expr; $r:expr, $c:expr) => (
        vec![vec![$elem; $c]; $r]
    );
    ($($x:expr),*) => (
        [$($x),*].iter().map(|v| v.to_vec()).collect::<Vec<_>>()
    );
    ($($x:expr,)*) => (mat![$($x),*])
}

impl Solution {
    pub fn reorganize_string(s: String) -> String {
        let n = s.len();
        let mut counts = (0..26)
            .map(|e| (e + b'a' as u8, 0))
            .collect::<Vec<(u8, usize)>>();
        s.bytes().into_iter().for_each(|e| {
            counts[(e - b'a') as usize].1 += 1;
        });
        counts.sort_by_key(|e| -(e.1 as i32));
        let half = (n + 1) / 2;
        if counts[0].1 > half {
            return String::from("");
        }

        let mut ret = Vec::new();
        ret.resize(n, 0_u8);
        let mut k = 0;
        for i in 0..half {
            if counts[k].1 == 0 {
                k += 1;
            }
            counts[k].1 -= 1;
            ret[i * 2] = counts[k].0;
        }
        for i in 0..n - half {
            if counts[k].1 == 0 {
                k += 1;
            }
            counts[k].1 -= 1;
            ret[i * 2 + 1] = counts[k].0;
        }

        String::from_utf8(ret).unwrap()
    }
}

fn main() {
    let cases = [
        String::from("aab"),
        String::from("aaab"),
        String::from("vvvlo"),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::reorganize_string(c.clone());
        println!("result: {:?}", r);
    }
}
