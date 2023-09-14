// 33. Search in Rotated Sorted Array

struct Solution {}

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let (mut i, mut j, mut m) = (0_usize, nums.len()-1, 0_usize);
        while i<j {
            m = (i+j)/2;
            if nums[i]<=nums[m] && nums[m]>nums[j] {
                i = m+1;
            } else if nums[i]>=nums[m] && nums[m]<nums[j] {
                j = m;
            } else {
                break;
            }
        };

        let pivot = i;

        let (mut i, mut j) = (0_usize, nums.len()-1);
        while i<j {
            m = (i+j)/2;
            let mnum = nums[(pivot+m)%nums.len()];
            if target > mnum {
                i = m+1;
            } else {
                j = m;
            }
        }

        i = (pivot + i)%nums.len();
        if nums[i] != target {
            return -1;
        }
        i as i32
    }
}

fn main() {
    let cases = [
        (vec![4,5,6,7,0,1,2], 0),
        (vec![4,5,6,7,0,1,2], 0),
        (vec![4,5,6,7,0,1,2], 0),
        (vec![3, 1], 0),
    ];
    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::search(c.0.clone(), c.1);
        println!("result is {}", r);
    }
}
