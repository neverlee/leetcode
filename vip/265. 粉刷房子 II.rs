// 265. 粉刷房子 II

struct Solution {}

impl Solution {
    pub fn min_cost_ii(costs: Vec<Vec<i32>>) -> i32 {
        let mut b1 = costs[0].clone();
        let mut b2 = costs[0].clone();

        let (il, jl) = (costs.len(), costs[0].len());

        let (mut p1, mut p2) = (&mut b1[..], &mut b2[..]);
        for i in 1..il {
            let m1 = p1.iter().enumerate().
            min_by_key(|&(_,e)|e).map(|(idx, _)| idx).unwrap();
            let m2 = p1.iter().enumerate().filter(|&(idx, _)| idx != m1).
            min_by_key(|&(_,e)|e).map(|(idx, _)| idx).unwrap();

            for j in 0..jl {
                p2[j] = p1[m1] + costs[i][j];
            }
            p2[m1] = p1[m2] + costs[i][m1];
            std::mem::swap(&mut p1, &mut p2);
        }

        *p1.iter().min().unwrap()
    }
}

fn main() {
    let cases = [
        (vec![vec![1,5,3], vec![2,9,4]], 2),
    ];
    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::min_cost_ii(c.0.clone());
        println!("result: {:?}", r);
    }
}
