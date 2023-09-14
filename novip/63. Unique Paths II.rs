// 63. Unique Paths II

struct Solution {}


fn array2d<T:Copy>(r:usize, c:usize, d:T) -> Vec<Vec<T>> {
    let mut ret = Vec::with_capacity(r);
    ret.resize_with(r, || {let mut ac = Vec::with_capacity(c); ac.resize(c, d); ac });
    ret
}

impl Solution {
    pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {
        let (m, n) = (obstacle_grid.len(), obstacle_grid[0].len());

        let mut f = array2d(m+1, n+1, 0);
        f[0][1] = 1;
        for i in 1..m+1 {
            for j in 1..n+1 {
                if obstacle_grid[i-1][j-1] == 0 {
                    f[i][j] += f[i-1][j];
                    f[i][j] += f[i][j-1];
                }
            }
        }
        f[m][n]
    }
}

fn main() {
    let cases = [
        (vec![vec![0,0,0],vec![0,1,0]], 0),
        (vec![vec![1]], 0),
    ];
    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::unique_paths_with_obstacles(c.0.clone());
        println!("result: {:?}", r);
    }
}

