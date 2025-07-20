mod apply_operations_to_an_array;

fn main() {    
    let result = apply_operations_to_an_array::apply_operations((&[847,847,0,0,0,399,416,416,879,879,206,206,206,272]).to_vec());
    println!("Result: {:?}", result);
}
