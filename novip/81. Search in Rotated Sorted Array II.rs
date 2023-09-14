// 81. Search in Rotated Sorted Array II

struct Solution {}

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> bool {
        let n = nums.len();
        if n == 0 {
            return false;
        }

        let (mut start, mut end, mut mid) = (0_usize, nums.len()-1, 0_usize);
        while start <= end {
            mid = (start + end)/2;
            if nums[mid] == target {
                return true;
            }

            if nums[start] == nums[mid] {
                start+=1;
                continue;
            }

            let pa = nums[start] <= nums[mid];
            let ta = nums[start] <= target;
            if pa ^ ta {
                if pa {
                    start = mid+1;
                } else {
                    end = mid-1;
                }
            } else {
                if nums[mid] < target {
                    start = mid+1;
                } else {
                    end = mid-1;
                }
            }
        }

        false
    }
}

fn main() {
    let cases = [
        (vec![4,5,6,7,0,1,2], 0),
        (vec![4,5,6,7,0,1,2], 3),
        (vec![4,5,6,7,0,1,2], 6),
        (vec![3, 1], 1),
        (vec![2,5,6,0,0,1,2], 0),
        (vec![2,5,6,0,0,1,2], 1),

    ];
    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::search(c.0.clone(), c.1);
        println!("result: {:?}", r);
    }
}
