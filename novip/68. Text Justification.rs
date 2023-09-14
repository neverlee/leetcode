// 68. Text Justification

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
    pub fn full_justify(words: Vec<String>, max_width: i32) -> Vec<String> {
        // pub fn full_justify(words: Vec<&str>, max_width: i32) -> Vec<String> {
        let mut cur = 0;
        let max_width = max_width as usize;

        let mut ret = Vec::new();
        while cur < words.len() {
            let mut len = 0;
            let mut lsum = 0;
            while cur + len < words.len() {
                if lsum + words[cur + len].len() + len > max_width {
                    break;
                }
                lsum += words[cur + len].len();
                len += 1;
            }
            let mut astr = String::with_capacity(max_width);
            let mut space = max_width - lsum;
            if cur + len == words.len() {
                for i in 0..len {
                    if i > 0 {
                        astr.push(' ');
                        space -= 1;
                    }
                    astr.push_str(&words[cur + i]);
                }
            } else {
                for i in 0..len {
                    if i > 0 {
                        let ns = (space + (len - i - 1)) / (len - i);
                        space -= ns;
                        for _ in 0..ns {
                            astr.push(' ');
                        }
                    }
                    astr.push_str(&words[cur + i]);
                }
            }
            for _ in 0..(space) {
                astr.push(' ');
            }
            ret.push(astr);
            cur += len;
        }

        ret
    }
}

fn main() {
    let cases = [
        (
            vecs![
                "This",
                "is",
                "an",
                "example",
                "of",
                "text",
                "justification.",
            ],
            16,
        ),
        (
            vecs!["What", "must", "be", "acknowledgment", "shall", "be"],
            16,
        ),
        (
            vecs![
                "Science",
                "is",
                "what",
                "we",
                "understand",
                "well",
                "enough",
                "to",
                "explain",
                "to",
                "a",
                "computer.",
                "Art",
                "is",
                "everything",
                "else",
                "we",
                "do",
            ],
            20,
        ),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::full_justify(c.0.clone(), c.1);
        for x in &r {
            println!("{}_", x);
        }
        println!("result: {:?}", r);
    }
}
