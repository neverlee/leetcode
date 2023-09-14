// 2814. Minimum Time Takes to Reach Destination Without Drowning

struct Solution {}

macro_rules! tensor {
    ($elem:expr; $($x:expr), +) => (
        // $crate::vec::from_elem($elem, $n)
        let shapes = [$($x),+];
        // for shape in shapes.iter()
        println!(">>>>>> {} {}", stringify!($elem), stringify!([$($x),+]) );
    );
    ($($x:expr),*) => (
        <[_]>::into_vec(box [$($x),*])
    );
    ($($x:expr,)*) => ($crate::vec![$($x),*])
}

macro_rules! mat {
    ($elem:expr; $r:expr, $c:expr) => (
        vec![vec![$elem; $c]; $r]
    );
    ($($x:expr),*) => (
        // <[_]>::into_vec(box [$($x),*])
        [$($x),*].iter().map(|v| v.to_vec()).collect::<Vec<_>>()
    );
    ($($x:expr,)*) => (mat[$($x),*])
}

impl Solution {
    pub fn minimum_seconds(land: Vec<Vec<String>>) -> i32 {
        // pub fn minimum_seconds(land: Vec<Vec<&str>>) -> i32 {
        let (m, n) = (land.len(), land[0].len());
        let mut fd = mat![-1; m, n];

        let mut q1 = Vec::new();
        let mut q2 = Vec::new();

        let mut bland = mat![b'X'; m+2, n+2];
        let mut start = (0, 0);
        let mut dist = (0, 0);
        q1.push((0, 0, 0));

        for (i, row) in land.iter().enumerate() {
            for (j, e) in row.iter().enumerate() {
                let c = e.as_bytes()[0];
                let (ti, tj) = (i + 1, j + 1);
                bland[ti][tj] = c;
                match c {
                    b'S' => start = (ti as i32, tj as i32),
                    b'D' => dist = (ti as i32, tj as i32),
                    b'*' => q1.push((ti as i32, tj as i32, -1)),
                    _ => (),
                }
            }
        }

        let dirs = [(-1, 0), (1, 0), (0, -1), (0, 1)];

        q1[0] = (start.0, start.1, 0);
        let (mut pq1, mut pq2) = (&mut q1, &mut q2);
        while pq1.len() > 0 {
            for &(x, y, t) in pq1.iter() {
                for dir in dirs {
                    let (nx, ny) = (x + dir.0, y + dir.1);
                    let nb = bland[nx as usize][ny as usize];
                    if nb == b'X' {
                        continue;
                    }
                    if t >= 0 && bland[x as usize][y as usize] != b'*' && (nb == b'.' || nb == b'D')
                    {
                        if nb == b'D' {
                            return t + 1;
                        }
                        bland[nx as usize][ny as usize] = b'S';
                        pq2.push((nx, ny, t + 1));
                    }

                    if t < 0 && (nb == b'.' || nb == b'S') {
                        bland[nx as usize][ny as usize] = b'*';
                        pq2.push((nx, ny, t - 1));
                    }
                }
            }

            pq1.clear();

            std::mem::swap(&mut pq1, &mut pq2);
        }

        -1
    }
}

fn main() {
    let cases = [
        mat![["D", ".", "*"], [".", ".", "."], [".", "S", "."]],
        mat![["D", "X", "*"], [".", ".", "."], [".", ".", "S"]],
        mat![
            ["D", ".", ".", ".", "*", "."],
            [".", "X", ".", "X", ".", "."],
            [".", ".", ".", ".", "S", "."]
        ],
        mat![["X", "X", "."], ["D", "X", "S"], [".", ".", "X"]],
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::minimum_seconds(c.clone());
        println!("result: {:?}", r);
    }
}
