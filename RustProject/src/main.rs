mod convert_binary_number_in_a_linked_list;


use convert_binary_number_in_a_linked_list::{Solution, ListNode};


fn build_linked_list(values: Vec<i32>) -> Option<Box<ListNode>> {
    let mut head: Option<Box<ListNode>> = None;

    for &value in values.iter().rev() {
        head = Some(Box::new(ListNode {
            val: value,
            next: head,
        }));
    }

    head
}

fn main() {
    let binary_values = vec![1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0];

    let head = build_linked_list(binary_values);

    let result = Solution::get_decimal_value(head);

    println!("Result: {}", result);
}
