mod count_hills_and_valleys_in_an_array;

use count_hills_and_valleys_in_an_array::{Solution};

fn main() {
    let nums = vec![2,4,1,1,6,5]
    let result = Solution::count_hill_valley(nums);

    println!("{:?}", result);
}


/*
pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
   let mut result: Vec<Vec<i32>> = Vec::new();
   let mut res: Vec<i32> = Vec::new(); 
   for i in 1..=num_rows {
      for j in 0..i {
         res.push(num_rows);
      }
      result.push(res.clone());
      res.clear();
   }

   result
}
*/