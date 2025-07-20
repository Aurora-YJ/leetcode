mod valid_word;

fn main() {    
    let result = valid_word::is_valid(("UuE6").to_string());
    println!("Result: {}", result);
}
