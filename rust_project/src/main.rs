mod count_hills_and_valleys_in_an_array;

use count_hills_and_valleys_in_an_array::{Solution};

fn main() {
    let nums = vec![2,4,1,1,6,5]
    let result = Solution::count_hill_valley(nums);

    println!("{:?}", result);
}
