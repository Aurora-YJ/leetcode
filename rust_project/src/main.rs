// mod sum_triplet;

// use sum_triplet::{Solution};

// fn main() {
//     let nums = vec![-1,0,1,2,-1,-4];



//     let result = Solution::three_sum(nums);

//     println!("{:?}", result);
// }

mod maximum_score_from_removing_substrings;

use maximum_score_from_removing_substrings::{Solution};

fn main(){
    let (x , y) = (4 , 5);

    let result = Solution::maximum_gain("cdbcbbaaabab".to_string(), x, y);
    println!("{:?}", result);
}