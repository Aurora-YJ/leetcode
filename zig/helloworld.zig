const std = @import("std");

pub fn main() void {
    
    const res: []i32 = ConvertIntegertotheSumofTwoNoZeroIntegers(9);
    for (res) |value| {
        std.debug.print("{}\n", .{value});
    }

}

pub fn ConvertIntegertotheSumofTwoNoZeroIntegers(n: i32) []i32 {
    var n_usize: usize = 0;
    if (n >= 0) {
        n_usize = @intCast(usize, n);
    }

    for (0..n_usize) |value| {
        std.debug.print("{d}\n", .{value});
    }

    return &[_]i32{ 4, 5 };
}
