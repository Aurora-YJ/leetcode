pub struct Solution;

impl Solution {
    pub fn maximum_gain(s: String, x: i32, y: i32) -> i32 {

        let mut res = 0;

        let mut is_x = false;
        if x > y {
            is_x = true;
        }


        let mut stack: Vec<char> = Vec::new();

        let mut char1 = 'a';
        let mut char2 = 'b';

        if !is_x {
            char1 = 'b';
            char2 = 'a';   
        }

        println!("{}{}", char1, char2);
        

        for v in s.chars() {
            if v == char2 && stack.len() > 0 && stack[stack.len()-1] == char1 {
                if is_x {
                    res += x;
                } else {
                    res += y;
                }
                stack.pop();
                continue;
            }

            stack.push(v);  
        }

        let mut stack2: Vec<char> = Vec::new();

        for (i,v) in stack.iter().enumerate() {

            if *v == char1 && stack2.len() > 0 && stack2[stack2.len()-1] == char2 {
                if is_x {
                    res += y;
                } else {
                    res += x;
                }
                stack2.pop();
                continue;
            }

            stack2.push(*v);  


        }

        println!("{:?}", stack);
        println!("{:?}", stack2);

        res
    }
}