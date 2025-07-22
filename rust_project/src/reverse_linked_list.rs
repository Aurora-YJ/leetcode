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
         next: None,
       }));



       let mut c = head.as_ref();
       
       while let Some(nd) = c {

          if let Some(ref mut inner_node) = node {
            
            if nd.val == left {
              inner_node.val = nd.val;
              inner_node.next = nd.next.clone(); 
            }

            if nd.val == right {
              inner_node.val = nd.val;
              inner_node.next = nd.next.clone(); 
            }
          }

        

        


          c = nd.next.as_ref();
       }

       node
    }
}

