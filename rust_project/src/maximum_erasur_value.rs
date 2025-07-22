pub struct Solution;


//[4,2,4,5,6]
// 0,1,2,3,4 

// 16911


/*
Input
nums =
[10000,1,10000,1,1,1,1,1,1]

Use Testcase
Output
10000
Expected
10001
*/

impl Solution {
    pub fn maximum_unique_subarray(nums: Vec<i32>) -> i32 {


        let mut array: Vec<i32> = Vec::new();

        let mut res = 0;
        let mut result = 0;

        let mut i = 0;



        while i < nums.len() {
            if array.contains(&nums[i]) {        

                while array.contains(&nums[i]) {   
                    res -= array[0];
                    array.remove(0);    
                }

               

            } else {
                res += nums[i];
                array.push(nums[i]);
                i+=1;
            }

            if res > result {
               result = res;
            }
            
        }
     
        if res > result {
            result = res;
        }    
        result
    }
}
