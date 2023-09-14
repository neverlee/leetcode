// 225. Implement Stack using Queues

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


#[warn(unused_macros)]
macro_rules! ifv {
    ($c:expr, $a:expr, $b: expr) => {
        if $c {
            $a
        } else {
            $b
        }
    };
}

use std::cell::RefCell;
use std::fmt;
use std::rc::{Rc, Weak};

struct Node<T: Copy + fmt::Display> {
    data: T,
    prev: Option<Rc<RefCell<Node<T>>>>,
    next: Option<Rc<RefCell<Node<T>>>>,
}

struct Queue<T>
where
    T: Copy + fmt::Display,
{
    front: Option<Rc<RefCell<Node<T>>>>,
    back: Option<Rc<RefCell<Node<T>>>>,
    size: usize,
}

impl<T: Copy + fmt::Display> Queue<T> {
    fn new() -> Queue<T> {
        Queue {
            front: None,
            back: None,
            size: 0,
        }
    }

    fn push(&mut self, e: T) {
        self.size += 1;
        let nnode = Node {
            data: e,
            prev: self.back.clone(),
            next: None,
        };
        let back = self.back.take();
        if let Some(nback) = back {
            let mut mnb = nback.borrow_mut();
            mnb.next = Some(Rc::new(RefCell::new(nnode)));
            self.back = mnb.next.clone();
        } else {
            self.front = Some(Rc::new(RefCell::new(nnode)));
            self.back = self.front.clone();
        }
    }
    fn pop(&mut self) -> Option<T> {
        let front = self.front.take();
        if let Some(nfront) = front {
            let mut cur = nfront.borrow_mut();
            let ret = Some(cur.data);
            let snext = cur.next.take();
            self.front = snext.clone();
            if let Some(next) = snext {
                next.borrow_mut().prev.take();
            } else {
                self.back.take();
            }
            self.size -= 1;
            return ret;
        }
        None
    }
    fn peek(&self) -> Option<T> {
        if let Some(n) = self.front.clone() {
            return Some(n.borrow().data);
        }
        None
    }
    fn len(&self) -> usize {
        self.size
    }
}

impl<T: Copy + fmt::Display> fmt::Display for Queue<T> {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "output list:\n .   size: {}", self.len())?;
        write!(f, "\n .   front: ")?;
        let mut p = self.front.clone();
        while let Some(pn) = p {
            let bpn = pn.borrow();
            write!(f, "{}-> ", bpn.data)?;
            p = bpn.next.clone();
        }
        write!(f, "\n .   back : ")?;
        let mut p = self.back.clone();
        while let Some(pn) = p {
            let bpn = pn.borrow();
            write!(f, "{}-> ", bpn.data)?;
            p = bpn.prev.clone();
        }
        writeln!(f, "")
    }
}

struct MyStack {
    q1: Queue<i32>,
    q2: Queue<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyStack {
    fn new() -> Self {
        MyStack {
            q1: Queue::new(),
            q2: Queue::new(),
        }
    }

    fn push(&mut self, x: i32) {
        self.q1.push(x);
        for _ in 1..self.q1.len() {
            if let Some(e) = self.q1.pop() {
                self.q1.push(e);
            }
        }
    }

    fn pop(&mut self) -> i32 {
        self.q1.pop().unwrap()
    }

    fn top(&self) -> i32 {
        self.q1.peek().unwrap()
    }

    fn empty(&self) -> bool {
        self.q1.size == 0
    }
}

/**
 * Your MyStack object will be instantiated and called as such:
 * let obj = MyStack::new();
 * obj.push(x);
 * let ret_2: i32 = obj.pop();
 * let ret_3: i32 = obj.top();
 * let ret_4: bool = obj.empty();
 */

fn main() {
    let cases = [
        vec![0, 1, 3, 5, 6, 8, 12, 17],
        // vec![0, 1, 2, 3, 4, 8, 9, 11],
    ];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);

        let mut q = Queue::new();
        for e in c.clone() {
            q.push(e);
            println!(">> push {}", q);
        }
        while let Some(e) = q.pop() {
            println!(">> pop {} {:?} {}", e, q.peek(), q);
        }

        // let r = Solution::can_cross(c.clone());
        // println!("result: {:?}", r);
    }
}
