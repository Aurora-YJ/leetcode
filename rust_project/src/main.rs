mod sum_triplet;

use sum_triplet::{Solution};

fn main() {
    let nums = vec![-1,0,1,2,-1,-4];



    let result = Solution::three_sum(nums);

    println!("{:?}", result);
} 