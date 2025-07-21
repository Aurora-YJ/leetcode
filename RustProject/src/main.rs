// mod reverse_linked_list;


// use reverse_linked_list::{Solution, ListNode};


// fn build_linked_list(values: Vec<i32>) -> Option<Box<ListNode>> {
//     let mut head: Option<Box<ListNode>> = None;

//     for &value in values.iter().rev() {
//         head = Some(Box::new(ListNode {
//             val: value,
//             next: head,
//         }));
//     }

//     head
// }

// fn main() {
//     let binary_values = vec![1,2,3,4,5];

//     let head = build_linked_list(binary_values);

//     let left = 2;
//     let right = 4;

//     let result = Solution::reverse_between(head, left , right);

//     println!("Result: {:?}", result);
// }


mod delete_characters_to_make_fancy_string;

use delete_characters_to_make_fancy_string::{Solution};

fn main() {
    let result = Solution::make_fancy_string("leeeetcoooodeeee     aaaabaaaa".to_string());

    println!("{}", result);
} 