pub struct Solution;

impl Solution {
    pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut result: Vec<Vec<i32>> = Vec::new();

        for i in 0..nums.len() {
            for j in 0..nums.len() {
                for k in 0..nums.len() {
                    if i != j && j != k && i != k {
                        if nums[i] + nums[j] + nums[k] == 0 {

                            if Self::is_checked(result.clone(),  nums[i], nums[j], nums[k]) {
                                result.push(vec![nums[i], nums[j], nums[k]]);
                            }
                        } 
                    }
                }
            }
        }


        result
    }


   fn is_checked(res: Vec<Vec<i32>>, i: i32, j: i32, k: i32) -> bool {
        if res.len() == 0 {
            return true;
        }

        let mut new_arr = vec![i, j, k];

        new_arr.sort();

        for arr in res {

            let mut sort_arr = arr.clone();

            sort_arr.sort();

            if sort_arr == new_arr {
                return false;
            }

        }

        true
    }




}