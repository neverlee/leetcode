// 518. Coin Change II
struct Solution {}

impl Solution {
    pub fn change(amount: i32, coins: Vec<i32>) -> i32 {
        let mut f = Vec::new();
        let n = amount as usize;
        f.resize(n as usize + 1, 0);

        f[0] = 1;

        for c in coins {
            for i in c as usize..n+1 {
                f[i] = f[i] + f[i-c as usize];
            }
        }

        f[n]
    }
}

fn main() {
    let cases = [
        (vec![1,2,5], 5),
        (vec![2], 3),
    ];
    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::change(c.1, c.0.clone());
        println!("result: {:?}", r);
    }
}
