// 847. Shortest Path Visiting All Nodes

struct Solution {}

#[warn(unused_macros)]
macro_rules! mat {
    ($elem:expr; $r:expr, $c:expr) => (
        vec![vec![$elem; $c]; $r]
    );
    ($($x:expr),*) => (
        [$($x),*].iter().map(|v| v.to_vec()).collect::<Vec<_>>()
    );
    ($($x:expr,)*) => (mat![$($x),*])
}

#[warn(unused_macros)]
macro_rules! vecs {
    ($($x:expr),*) => (
        vec![$($x),*].into_iter().map(|s| s.to_string()).collect::<Vec<_>>()
    );
    ($($x:expr,)*) => (vecs![$($x),*])
}

impl Solution {
    pub fn shortest_path_length(graph: Vec<Vec<i32>>) -> i32 {
        let n = graph.len();

        let unlimit = 100;
        let mut fm = mat![unlimit; n, 1<<n];
        for i in 0..n {
            fm[i][1 << i] = 0;
        }

        loop {
            let mut update = false;
            for f in 0..n {
                for k in 0..(1 << n) {
                    if fm[f][k] == unlimit {
                        continue;
                    }

                    for &t in graph[f].iter() {
                        let t = t as usize;
                        let nk = k | (1 << t);
                        if fm[t][nk] > fm[f][k] + 1 {
                            update = true;
                            fm[t][nk] = fm[f][k] + 1;
                        }
                    }
                }
            }
            if !update {
                break;
            }
        }

        fm.iter()
            .map(|v| v[(1 << n) as usize - 1])
            .fold(unlimit, |a, b| a.min(b))
    }
}

fn main() {
    let cases = [
        (vec![vec![1, 2, 3], vec![0], vec![0], vec![0]]),
        (vec![vec![1], vec![0, 2, 4], vec![1, 3, 4], vec![2], vec![1, 2]]),
        (vec![
            vec![1],
            vec![0, 2, 4],
            vec![1, 3],
            vec![2],
            vec![1, 5],
            vec![4],
        ]),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let r = Solution::shortest_path_length(c.clone());
        println!("result: {:?}", r);
    }
}
