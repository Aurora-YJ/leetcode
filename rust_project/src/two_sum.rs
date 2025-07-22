pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut arr: Vec<i32> = Vec::new();
        for i in 0..nums.len(){
             for j in 0..nums.len(){
                if i != j && nums[i] + nums[j] == target{
                    arr.push(i as i32);
                    arr.push(j as i32);
                    return arr;
                }
            }
        }
    arr
}
