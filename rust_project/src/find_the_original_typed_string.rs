use std::collections::{HashMap};


/*

Input
word =
"ere"

Use Testcase
Output
2
Expected
1

*/
pub fn possible_string_count(word: String) -> i32 {

    let mut mapp = HashMap::new();
    let mut lastl: Option<char> = None;
    for v in word.chars() {
        if Some(v) == lastl {
            *mapp.entry(v).or_insert(0) += 1;
        } else {
            mapp.entry(v).or_insert(1);
        }
        lastl = Some(v);
    }
    
    let mut c = 1;

    for (_,val) in &mapp {
        if *val > 1 {
            c = c + (val - 1);
        }
    }
    c
}
