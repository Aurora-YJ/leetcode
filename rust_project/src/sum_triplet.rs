pub struct Solution;

impl Solution {
    pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
        
        let mut res: Vec<i32> = Vec::new();
        let mut result: Vec<Vec<i32>> = Vec::new();

        let mut cont = 0;

        for i in &nums {
            cont += i;
            res.push(*i);
            if cont == 0 {
                println!("{}", cont);
                break;
            }
        }

        result.push(res);
        result
    }
}