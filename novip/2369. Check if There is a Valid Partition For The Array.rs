// 2369. Check if There is a Valid Partition For The Array

struct Solution {}


impl Solution {
    pub fn valid_partition(nums: Vec<i32>) -> bool {
        let n = nums.len();

        let mut f = Vec::new();
        f.resize(nums.len(), false);

        if n >= 2 && nums[0] == nums[1] {
            f[1] = true;
        }
        if n >= 3 {
            let (a, b) = (nums[1] - nums[0], nums[2] - nums[1]);
            if a == b && (a == 0 || a == 1) {
                f[2] = true;
            };
        }

        for i in 3..n {
            let (a, b) = (nums[i] - nums[i-1], nums[i-1] - nums[i-2]);
            f[i] = ( f[i-2] && a == 0 ) || 
                ( f[i-3] && a == b && (a == 0 || a == 1) );
        }

        f[n-1]
    }
}

fn main() {
    let cases = [
        (vec![4,4,4,5,6]),
        (vec![1,1,1,2]),
        (vec![1,1,1]),
        (vec![1,1]),
        (vec![1,2]),
        (vec![993335,993336,993337,993338,993339,993340,993341]),
    ];
    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::valid_partition(c.clone());
        println!(">>>>>>>>> {}", r);
    }
}