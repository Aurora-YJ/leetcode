pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

pub struct Solution;

impl Solution {
    pub fn get_decimal_value(head: Option<Box<ListNode>>) -> i32 {
        let  mut result = 0;

        let mut currant = head.as_ref();
        while let Some(node) = currant {
            result <<= 1;
            result = result + node.val;
            currant = node.next.as_ref(); 
        }

        result
    }
}