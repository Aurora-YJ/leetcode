pub struct Solution ;

impl Solution {
    pub fn make_fancy_string(s: String) -> String {
        let mut help: i32 = 0;
        let mut the_dgt: char = '\0';

        let mut res: Vec<char> = Vec::new();

        let new_str: Vec<char> = s.chars().collect();
        for   (i, ch) in new_str.iter().enumerate()  {


            if the_dgt != *ch {
                the_dgt =*ch;
                help = 0;
            }

            if  i < new_str.len()-1  && new_str[i] == new_str[i+1] || new_str[i] == the_dgt {
                help+=1;
            }

            if help > 2 {

                continue;
            }

            the_dgt = *ch;

            res.push(*ch);

        }
        let n = res.into_iter().collect();
        n
    }
}