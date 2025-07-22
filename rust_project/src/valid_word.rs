pub fn is_valid(word: String) -> bool {

    if word.len() < 3 {
        return false;
    }

    let mut if_one_vowel = false ;
    let mut if_one_consonant = false;

    let mut helpp = false;
    for i in  word.chars() {
        helpp = false;
        if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' {
            if !is_vowel(i) {
                if_one_consonant = true;
                helpp = true;
            } else {
                if_one_vowel = true;
                helpp = true;
            }
        }

        if i >= '0' && i <= '9' {
            helpp = true;
        }

        if !helpp {
            return false;
        }
    }
    

   
    
    if_one_vowel && if_one_consonant
}


fn is_vowel(v: char)->  bool {
    
    let vowels: [char; 10] = ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'];

    for ch in &vowels {
        if v == *ch {
            return true;
        }
    }

    false
}