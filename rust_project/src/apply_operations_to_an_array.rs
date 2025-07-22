pub fn apply_operations(nums: Vec<i32>) -> Vec<i32> {
    let mut res: Vec<i32> = nums;
    for i in 0..res.len()-1 {
        if res[i] == res[i+1] {
            res[i] = res[i] + res[i+1];
            res[i+1] = 0;
        }
    }

    println!("{:?}", res);

    let mut helper = true;
    while helper {
        helper = false;
        for i in 0..res.len()-1 {
        if res[i] == 0 && res[i+1] > 0 {
            helper = true;
            let t = res[i+1];
            res[i+1] = res[i];
            res[i] = t;
        }
    }
    }
    
    res
}
/*
[847,847,0,0,0,399,416,416,879,879,206,206,206,272]

[1694, 0,0,0,0,399, 832, 0, 1785, 0, 412, 0, 206, 272]

[1694,399,832,1758,412,206,272,0,0,0,0,0,0,0]
*/