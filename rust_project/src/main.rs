mod count_hills_and_valleys_in_an_array;

use count_hills_and_valleys_in_an_array::{Solution};

fn main() {
    let nums = vec![6,6,5,5,4,1];

    let result = Solution::count_hill_valley(nums);

    println!("{:?}", result);
}
