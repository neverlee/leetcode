// 86. Partition List

struct Solution {}

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }

    fn from(ay: &[i32]) -> Option<Box<ListNode>> {
        let mut head = ListNode::new(0);
        for a in ay.iter().rev() {
            let node = ListNode {
                val: *a,
                next: head.next.take(),
            };
            head.next = Some(Box::new(node));
        }
        head.next
    }

    fn output(head: &Option<Box<ListNode>>) {
        let mut p = head;
        print!("list is: ");
        while let Some(cur) = p {
            print!("{},", cur.val);
            p = &cur.next;
        }
        println!("");
    }
}

impl Solution {
    pub fn partition(head: Option<Box<ListNode>>, x: i32) -> Option<Box<ListNode>> {
        let mut p = &head;
        let mut less: Option<Box<ListNode>> = None;
        let mut more: Option<Box<ListNode>> = None;

        let mut pless = &mut less as *mut Option<_>;
        let mut pmore = &mut more as *mut Option<_>;
        while let Some(cur) = p {
            p = &cur.next;
            if cur.val < x {
                let mut node = Box::new(ListNode::new(cur.val));
                let pnext = &mut node.next as *mut Option<_>;
                unsafe {
                    *pless = Some(node);
                }
                pless = pnext;
            } else {
                let mut node = Box::new(ListNode::new(cur.val));
                let pnext = &mut node.next as *mut Option<_>;
                unsafe {
                    *pmore = Some(node);
                }
                pmore = pnext;
            }
        }
        unsafe {
            *pless = more;
        }
        // ListNode::output(&less);
        less
    }
}

fn main() {
    let cases = [(vec![1, 4, 3, 2, 5, 2], 3), (vec![2, 1], 2)];

    for (i, c) in cases.iter().enumerate() {
        println!("start test case {}: {:?}", i, c);
        let list = ListNode::from(&c.0);
        let r = Solution::partition(list, c.1);
        println!("result: {:?}", r);
    }
}
