// 542. 01 Matrix

struct Solution {}

macro_rules! mat {
    ($elem:expr; $r:expr, $c:expr) => (
        vec![vec![$elem; $c]; $r]
    );
    ($($x:expr),*) => (
        [$($x),*].iter().map(|v| v.to_vec()).collect::<Vec<_>>()
    );
    ($($x:expr,)*) => (mat![$($x),*])
}

impl Solution {
    pub fn update_matrix(mat: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let (m, n) = (mat.len(), mat[0].len());
        let mut ret = mat![0; m, n];

        let mut q1 = Vec::new();
        let mut q2 = Vec::new();

        let dirs = [(-1, 0), (1, 0), (0, -1), (0, 1)];

        for (i, v) in mat.iter().enumerate() {
            for (j, e) in v.iter().enumerate() {
                if *e == 0 {
                    q1.push((i as i32, j as i32));
                }
            }
        }

        let (mut pq1, mut pq2) = (&mut q1, &mut q2);
        let mut inx = 1;
        while pq1.len() > 0 {
            for &(x, y) in pq1.iter() {
                for dir in dirs {
                    let (nx, ny) = (x + dir.0, y + dir.1);
                    let (nxs, nys) = (nx as usize, ny as usize);
                    if nxs >= 0 && nxs < m && nys >= 0 && nys < n {
                        if mat[nxs][nys] == 1 && ret[nxs][nys] == 0 {
                            ret[nxs][nys] = inx;
                            pq2.push((nx, ny));
                        }
                    }
                }
            }

            inx += 1;
            pq1.clear();
            std::mem::swap(&mut pq1, &mut pq2);
        }

        ret
    }
}

fn main() {
    let cases = [mat![[0, 0, 0], [0, 1, 0], [1, 1, 1]]];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::update_matrix(c.clone());
        println!("result: {:?}", r);
    }
}
