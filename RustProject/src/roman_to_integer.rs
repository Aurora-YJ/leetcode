pub fn roman_to_integer(s: String) -> i32 {

    fn value(c: char)  -> i32 {
        match c {
            'I' => 1,
            'V' => 5,
            'X' => 10,
            'L' => 50,
            'C' => 100,
            'D' => 500,
            'M' => 1000,
            _=> 0,
        }
    }

   let  mut res : i32 = 0;
   let chars : Vec<char> = s.chars().collect();
   let  mut i :usize = 0;

    while i < chars.len() {
        let rn1 = value(chars[i]);
        
            if i+1 != chars.len() {

                let rn2 = value(chars[i+1]);
                if rn1 < rn2 {
                        res +=  rn2 - rn1;
                       i += 2;
    
                       continue;
                }
                  

            }


        res += rn1;
        i += 1; 
    }    
    res
}

