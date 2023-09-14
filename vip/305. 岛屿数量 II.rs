// 305. 岛屿数量 II

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

#[derive(Clone, PartialEq, PartialOrd, Debug)]
enum Block {
    W,
    L(usize, usize),
}

#[derive(Debug)]
struct Island {
    m: Vec<Vec<Block>>,
}

impl Island {
    fn new(m: usize, n: usize) -> Island {
        Island {
            m: mat![Block::W; m, n],
        }
    }

    fn set_land(&mut self, x: usize, y: usize) -> bool {
        if self.m[x][y] == Block::W {
            self.m[x][y] = Block::L(x, y);
            return true;
        }
        return false;
    }

    fn find(&mut self, x: usize, y: usize) -> Block {
        if let Block::L(nx, ny) = self.m[x][y] {
            if (nx, ny) == (x, y) {
                return Block::L(nx, ny);
            } else {
                let r = self.find(nx, ny);
                self.m[x][y] = r.clone();
                return r;
            }
        }

        Block::W
    }

    fn merge(&mut self, a: (usize, usize), b: (usize, usize)) -> bool {
        let ah = self.find(a.0, a.1);
        let bh = self.find(b.0, b.1);
        if ah != bh {
            // println!("----- merge {:?} {:?}         {:?} {:?}", a, ah, b, bh);
            if let (Block::L(ax, ay), Block::L(bx, by)) = (ah, bh) {
                self.m[ax][ay] = self.m[bx][by].clone();
                return true;
            }
        }
        false
    }
}

impl Solution {
    pub fn num_islands2(m: i32, n: i32, positions: Vec<Vec<i32>>) -> Vec<i32> {
        let mut island = Island::new((m + 2) as usize, (n + 2) as usize);
        let mut num = 0;
        let mut ret = Vec::with_capacity(positions.len());

        let dirs: [(i32, i32); 4] = [(0, -1), (0, 1), (-1, 0), (1, 0)];
        for p in positions {
            let (px, py) = (p[0] + 1, p[1] + 1);
            if island.set_land(px as usize, py as usize) {
                num += 1;
            }
            for d in dirs {
                let (nx, ny) = (px + d.0, py + d.1);
                if island.merge((px as usize, py as usize), (nx as usize, ny as usize)) {
                    num -= 1;
                }
            }

            ret.push(num);
        }
        ret
    }
}

fn main() {
    let cases = [
        (3, 3, mat![[0, 0], [0, 1], [1, 2], [2, 1]]),
        (1, 1, mat![[0, 0]]),
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let r = Solution::num_islands2(c.0.clone(), c.1.clone(), c.2.clone());
        println!("result: {:?}", r);
    }
}
