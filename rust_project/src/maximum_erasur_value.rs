use std::collections::{HashSet};

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

        let mut sett: HashSet<i32> = HashSet::new();

        let mut start = 0;
        let mut end = 0;

        let mut sum = 0;
        let mut max_sum = 0;

        while end < nums.len() {
            if let Some(_n) = sett.get(&nums[end]) {
                sett.remove(&nums[start]);
                sum -= nums[start];
                start += 1;
                
            } else {
                
                sett.insert(nums[end]);
                sum += nums[end];
                if sum > max_sum {
                    max_sum = sum;
                }
                end+= 1;
            }
               
        }
        max_sum
    }
}
