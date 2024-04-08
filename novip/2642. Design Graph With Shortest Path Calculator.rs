// 1793. Maximum Score of a Good Subarray
2642. Design Graph With Shortest Path Calculator

struct Solution {}

#[allow(unused_macros)]
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

// use std::{thread, time::Duration};
use std::cmp::{min, max};

struct Graph {
    n: usize,
    maze: Vec<Vec<i32>>,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl Graph {

    fn new(n: i32, edges: Vec<Vec<i32>>) -> Self {
        let nu = n as usize;
        let mut g = Graph{
            n: nu,
            maze: mat![i32::MAX; nu, nu],
        };
        for edge in edges.iter() {
            let (x, y, cost) = (edge[0], edge[1], edge[2]);
            let (xu, yu) = (x as usize, y as usize);
            g.maze[xu][yu] = cost;
        }

        for k in 0..nu {
            g.maze[k][k] = 0;
        }

        for k in 0..nu {
            for i in 0..nu {
                for j in 0..nu {
                    if g.maze[i][k] == i32::MAX || g.maze[k][j] == i32::MAX {
                        continue;
                    }
                    g.maze[i][j] = min(g.maze[i][j], g.maze[i][k] + g.maze[k][j]);
                }
            }
        }
        g
    }
    
    fn add_edge(&mut self, edge: Vec<i32>) {
        let (x, y, cost) = (edge[0], edge[1], edge[2]);
        let (xu, yu) = (x as usize, y as usize);

        if self.maze[xu][yu] < cost {
            return
        }

        self.maze[xu][yu] = cost;
        for i in 0..self.n {
            for j in 0..self.n {
                if self.maze[i][xu] == i32::MAX || self.maze[yu][j] == i32::MAX {
                    continue;
                }
                self.maze[i][j] = min(self.maze[i][j], self.maze[i][xu] + self.maze[yu][j] + cost);
            }
        }
    }
    
    fn shortest_path(&self, node1: i32, node2: i32) -> i32 {
        let r = self.maze[node1 as usize][node2 as usize];
        let r = if r == i32::MAX { -1 } else { r };
        r
    }
}

/**
 * Your Graph object will be instantiated and called as such:
 * let obj = Graph::new(n, edges);
 * obj.add_edge(edge);
 * let ret_2: i32 = obj.shortest_path(node1, node2);
 */



fn main() {
    let cases = [(vec![1,4,3,7,4,5], 3), (vec![5,5,4,5,4,1,1,1], 0)];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        // let r = Solution::maximum_score(c.0.clone(), c.1);
        // println!("result: {:?}", r);
    }
    let mut g = Graph::new(4, mat![[0,2,5],[0,1,2],[1,2,1],[3,0,3]]);
    g.shortest_path(3,2);
    g.shortest_path(0,3);
    g.add_edge(vec![1,3,4]);
    g.shortest_path(0,3);
}
