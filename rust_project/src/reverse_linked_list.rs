//Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]

pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}


impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

pub struct Solution;

impl Solution {
    pub fn reverse_between(head: Option<Box<ListNode>>, left: i32, right: i32) -> Option<Box<ListNode>> {
      
       let mut node =  Some(Box::new(ListNode {
         val : 0,
         next: head,
       }));



       let mut c = node.as_mut();
       

       
       while let Some(nd) = c {

          if nd.val == left {
            
          }

          if nd.val == right {
            
          }

          c = nd.next.as_mut();
       }

       node
    }
}

