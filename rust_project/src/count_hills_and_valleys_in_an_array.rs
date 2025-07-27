pub struct Solution;

impl Solution {
    pub fn count_hill_valley(nums: Vec<i32>) -> i32 {
      
       // vec![57,57,57,57,57,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,85,85,85,86,86,86];
       let mut count = 0;
        for i in 0..nums.len() {
            if i > 0 && i < nums.len()-1  {

                if i > 0 && nums[i] == nums[i - 1] {
                     continue;
                }
              
                let mut leftk = false;
                let mut lefts = false;
                
                let mut rightk = false;
                let mut rights = false;

                let mut c = 1;
                while i + c < nums.len() {

                    
                    if  nums[i] > nums[i+c] {
                        rights = true;

                        break;
                    } else if nums[i] < nums[i+c] {
                        rightk = true;


                        break;
                    } 

                   
                    c+=1;
                }



                let mut j = 1;

                while i >= j {
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