pub struct Solution;

impl Solution {
    pub fn count_hill_valley(nums: Vec<i32>) -> i32 {
      
       // [2,4,1,1,6,5]

       let mut count = 0;
        for i in 0..nums.len() {
            if i > 0 && i < nums.len()-1  {

                if nums[i] == nums[i-1] {
                    continue;
                }  
                let mut leftk = false;
                let mut lefts = false;
                
                let mut rightk = false;
                let mut rights = false;

                let mut c = 1;
                loop {
                    if  nums[i] > nums[i+c] {
                        rights = true;

                        break;
                    } else if nums[i] < nums[i+c] {
                        rightk = true;


                        break;
                    } 

                    if c == nums.len() -1 {
                        break;
                    }
                    c+=1;
                }



                let mut j = 1;
                loop {
                    
                    if  nums[i] > nums[i-j] {
                        lefts = true;
                        break;
                    } else if nums[i] < nums[i-j] {
                        leftk = true;
                        break;
                    } 
                    
                    j+=1;
                }


                if leftk && rightk {
                    count+=1;
                }

                if lefts && rights {
                    count+=1;
                }


            }
        }
        count
    }
}