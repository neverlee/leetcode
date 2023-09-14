// 499. 迷宫 III

struct Solution {}

macro_rules! mat {
    ($elem:expr; $r:expr, $c:expr) => (
        vec![vec![$elem; $c]; $r]
    );
    ($($x:expr),*) => (
        // <[_]>::into_vec(box [$($x),*])
        [$($x),*].iter().map(|v| v.to_vec()).collect::<Vec<_>>()
    );
    ($($x:expr,)*) => (mat![$($x),*])
}

impl Solution {
    pub fn find_shortest_way(maze: Vec<Vec<i32>>, ball: Vec<i32>, hole: Vec<i32>) -> String {
        let (m, n) = (maze.len(), maze[0].len());
        let mut bmaze = mat![1; m+2, n+2];
        let mut paths = mat![(0xfffffff, String::from("z") ); m+2, n+2];

        for (i, row) in maze.iter().enumerate() {
            for (j, e) in row.iter().enumerate() {
                let (ti, tj) = (i + 1, j + 1);
                bmaze[ti][tj] = maze[i][j];
            }
        }

        let dirs = [(-1, 0, 'u'), (1, 0, 'd'), (0, -1, 'l'), (0, 1, 'r')];
        let start = (ball[0] + 1, ball[1] + 1);
        let dist = (hole[0] + 1, hole[1] + 1);

        let mut queue = Vec::new();
        queue.push((start.0, start.1));
        paths[start.0 as usize][start.1 as usize] = (0, String::from(""));
        let mut front = 0;
        while front < queue.len() {
            let cur = queue[front];
            front += 1;

            for d in dirs {
                let mut len = 0;
                loop {
                    len += 1;
                    let next = (cur.0 + d.0 * len, cur.1 + d.1 * len);
                    if next == dist {
                        len = -len;
                        break;
                    }
                    if bmaze[next.0 as usize][next.1 as usize] == 1 {
                        break;
                    }
                }

                let mut npath = paths[cur.0 as usize][cur.1 as usize].clone();
                npath.1.push(d.2);
                match len {
                    0 | 1 => (),
                    l if l < 0 => {
                        npath.0 -= l;
                        if npath < paths[dist.0 as usize][dist.1 as usize] {
                            paths[dist.0 as usize][dist.1 as usize] = npath;
                        };
                    }
                    l => {
                        npath.0 += l - 1;
                        let next = (cur.0 + d.0 * (l - 1), cur.1 + d.1 * (l - 1));
                        if npath < paths[next.0 as usize][next.1 as usize] {
                            paths[next.0 as usize][next.1 as usize] = npath;
                            queue.push(next);
                        };
                    }
                }
            }
        }

        let mut ret = paths[dist.0 as usize][dist.1 as usize].clone().1;
        if ret == "z" {
            ret = String::from("impossible")
        }
        ret
    }
}

fn main() {
    let cases = [
        (
            mat![
                [0, 0, 0, 0, 0],
                [1, 1, 0, 0, 1],
                [0, 0, 0, 0, 0],
                [0, 1, 0, 0, 1],
                [0, 1, 0, 0, 0],
            ],
            vec![4, 3],
            vec![0, 1],
        ),
        (
            mat![
                [0, 0, 0, 0, 0],
                [1, 1, 0, 0, 1],
                [0, 0, 0, 0, 0],
                [0, 1, 0, 0, 1],
                [0, 1, 0, 0, 0],
            ],
            vec![4, 3],
            vec![3, 0],
        ),
        (
            mat![
                [0, 0, 0, 0, 0, 0, 0],
                [0, 0, 1, 0, 0, 1, 0],
                [0, 0, 0, 0, 1, 0, 0],
                [0, 0, 0, 0, 0, 0, 1],
            ],
            vec![0, 4],
            vec![3, 5],
        ),
        (
            mat![
                [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                [1, 1, 1, 1, 1, 1, 1, 1, 1, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 1, 1, 1, 1, 1, 1, 1, 1, 1],
                [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                [1, 1, 1, 1, 1, 1, 1, 1, 1, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 1, 1, 1, 1, 1, 1, 1, 1, 1],
                [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                [1, 1, 1, 1, 1, 1, 1, 1, 1, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 1, 1, 1, 1, 1, 1, 1, 1, 1],
                [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                [1, 1, 1, 1, 1, 1, 1, 1, 1, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 1, 1, 1, 1, 1, 1, 1, 1, 1],
                [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                [1, 1, 1, 1, 1, 1, 1, 1, 1, 0],
                [0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
                [0, 1, 1, 1, 1, 1, 1, 1, 1, 1],
            ],
            vec![0, 0],
            vec![19, 0],
        ),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::find_shortest_way(c.0.clone(), c.1.clone(), c.2.clone());
        println!("result: {:?}", r);
    }
}
