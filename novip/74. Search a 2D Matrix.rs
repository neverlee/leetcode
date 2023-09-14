// 74. Search a 2D Matrix

struct Solution {}

impl Solution {
    pub fn search_matrix(matrix: Vec<Vec<i32>>, target: i32) -> bool {
        let (m, n) = (matrix.len(), matrix[0].len());
        let end = m*n;

        let (mut i, mut j, mut mid) = (0, end, 0); 

        let get = |id:usize| matrix[id/n][id%n];

        while i<j {
            mid = (i+j)/2;
            if target > get(mid) {
                i = mid+1;
            } else {
                j = mid;
            }
        }
        i<end && get(i) == target
    }
}

fn main() {
    let cases = vec![3];
    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let matrix = vec![
            vec![1,3,5,7],vec![10,11,16,20],vec![23,30,34,60]
        ];
        let target = 13;
        Solution::search_matrix(matrix, target);
    }
}
