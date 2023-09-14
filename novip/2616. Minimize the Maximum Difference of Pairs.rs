// 2616. Minimize the Maximum Difference of Pairs

struct Solution {}

impl Solution {
    pub fn minimize_max(nums: Vec<i32>, p: i32) -> i32 {
        let mut nums = nums.clone();
        nums.sort();
        let (mut i, mut j, mut m) = (0, nums.last().unwrap()-nums[0], 0);
        while i<j {
            m = i+(j-i)/2;
            let cnt = Self::vaild_count(&nums[..], m);
            if cnt < p {
                i = m+1;
            } else {
                j = m;
            }
        }
        i
    }

    fn vaild_count(nums: &[i32], threshold: i32) -> i32 {
        let (mut i, mut c) = (1_usize, 0_i32);
        while i<nums.len() {
            if nums[i] - nums[i-1] <= threshold {
                c+=1; i+=1;
            }
            i+=1;
        }
        c
    }
}

fn main() {
    let cases = [
        (vec![1,1,0,3], 2),
        (vec![10,1,2,7,1,3], 2),
        (vec![4,2,1,2], 1),
    ];
    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        Solution::minimize_max(c.0.clone(), c.1);
    }
}
