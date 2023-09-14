// 239. Sliding Window Maximum

struct Solution {}

impl Solution {
    pub fn max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut ret = Vec::with_capacity(nums.len() + 1 - k as usize);
        let mut mq = Vec::with_capacity(nums.len());

        let k = k as usize;
        let mut mqat = 0;
        for i in 0..nums.len() {
            let n = nums[i];
            while mqat < mq.len() && *mq.last().unwrap() < n {
                mq.pop();
            }
            mq.push(n);

            if i >= k && mq[mqat] == nums[i - k] {
                mqat += 1;
            }

            if i >= k - 1 {
                ret.push(mq[mqat]);
            }
        }
        ret
    }
}

fn main() {
    let cases = [(vec![1, 3, -1, -3, 5, 3, 6, 7], 3), (vec![1], 1)];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::max_sliding_window(c.0.clone(), c.1);
        println!("result: {:?}", r);
    }
}
