const std = @import("std");

pub fn main() void {

    // const res: []i32 = ConvertIntegertotheSumofTwoNoZeroIntegers(69);
    ConvertIntegertotheSumofTwoNoZeroIntegers(11);
    // for (res) |value| {
    //     std.debug.print("{}\n", .{value});
    // }

}

pub fn ConvertIntegertotheSumofTwoNoZeroIntegers(n: i32) void {
    const y: usize = @intCast(n);
    for (1..y) |v| {
        const n1 = v;
        const n2 = y - n1;

        const h = myitoa(n1);
        std.debug.print("{any}\n", .{h});

        if (n1 + n2 == n) {
            
        }
    }
}

pub fn myitoa(n: usize) []u8 {
    var buffer: [20]u8 = undefined;
    var slice: []u8 = buffer[0..];

    var i: usize = 0;

    var num = n;
    while (num != 0) {
        const digit: u8 = @intCast(num % 10);
        slice[i] = digit + '0';

        num /= 10;
        i += 1;
    }
    

    return slice[0..i];
}
