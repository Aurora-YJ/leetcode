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


mod maximum_erasur_value;

use maximum_erasur_value::{Solution};

fn main() {
    // let nums = vec![187,470,25,436,538,809,441,167,477,110,275,133,666,345,411,459,490,266,987,965,429,166,809,340,467,
    // 318,125,165,809,610,31,585,970,306,42,189,169,743,78,810,70,382,367,490,
    // 787,670,476,278,775,673,299,19,893,817,971,458,409,886,434];

    //let nums = vec![4,2,4,5,6];


    let nums = vec![10000,1,10000,1,1,1,1,1,1];
    let result = Solution::maximum_unique_subarray(nums);

    println!("{}", result);
} 