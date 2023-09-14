// 215. Kth Largest Element in an Array

struct Solution {}

impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums.clone();
        let end = (nums.len() - 1) as i32;
        Self::ksearch(&mut nums, 0, end, k - 1)
    }

    fn ksearch(nums: &mut Vec<i32>, low: i32, high: i32, k: i32) -> i32 {
        if low < high {
            let pivot = nums[high as usize];
            let mut j = low - 1_i32;
            for i in low..high + 1 {
                if pivot <= nums[i as usize] {
                    j += 1;
                    nums.swap(i as usize, j as usize);
                }
            }

            if k < j {
                Self::ksearch(nums, low, j - 1, k)
            } else if k > j {
                Self::ksearch(nums, j + 1, high, k)
            } else {
                nums[j as usize]
            }
        } else {
            nums[low as usize]
        }
    }
}

fn main() {
    let cases = [
        (vec![3, 2, 1, 5, 6, 4], 2),
        (vec![3, 2, 3, 1, 2, 4, 5, 5, 6], 4),
        (vec![3], 1),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::find_kth_largest(c.0.clone(), c.1);
        println!("result {}", r);
    }
}
