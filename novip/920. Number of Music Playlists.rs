// 920. Number of Music Playlists
struct Solution {}

use std::cmp::min;
impl Solution {
    pub fn num_music_playlists(n: i32, goal: i32, k: i32) -> i32 {
        let (n, g, k) = (n as usize, goal as usize, k as usize);
        let m = 1e9 as i64 + 7;

        let mut dp = [[0_i64;101]; 101];
        dp[0][0] = 1;
        for i in 1..g+1 {
            for j in 1..min(i, n)+1 {
                dp[i][j] = dp[i-1][j-1] * (n-j+1) as i64 % m;
                if j>= k {
                    dp[i][j] = (dp[i][j] + dp[i-1][j]*((j-k) as i64))%m;
                }

            }
        }

        dp[g][n] as i32
    }
}


fn main() {
    let cases = vec![(3,3,1)];
    for (i, c) in cases.iter().enumerate() {
        let (n, g, k) = c;
        println!("start test case {}: {:?}", i, c);
        Solution::num_music_playlists(*n, *g, *k);
    }
}
