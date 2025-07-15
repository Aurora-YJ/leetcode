mod roman_to_integer;

fn main() {    
    let result = roman_to_integer::roman_to_integer("III".to_string());
    println!("Result: {}", result);
}
