mod sum_triplet;

use sum_triplet::{Solution};

fn main() {
    let nums = vec![-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0];



    let result = Solution::three_sum(nums);

    println!("{:?}", result);
}
