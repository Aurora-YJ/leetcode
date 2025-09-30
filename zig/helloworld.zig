const std = @import("std");

pub fn main() void {

    // const res: []i32 = ConvertIntegertotheSumofTwoNoZeroIntegers(69);
    ConvertIntegertotheSumofTwoNoZeroIntegers(3);
    // for (res) |value| {
    //     std.debug.print("{}\n", .{value});
    // }

}

pub fn ConvertIntegertotheSumofTwoNoZeroIntegers(n: i32) void {
    const y: usize = @intCast(n);
    for (1..y) |v| {
        const n1 = v;
        const n2 = y - n1;

        const str1 = myitoa(n1);
    const str2 = myitoa(n2);
    
    std.debug.print("n1: {s}, n2: {s}\n", .{str1, str2});

        if (n1 + n2 == n) {
            
        }
    }
}



pub fn myitoa(n: usize) []u8 {
    var buffer: [20]u8 = undefined;
    var i: usize = 20;

    std.debug.print("Converting: {}\n", .{n});
    
    var num = n;
    if (num == 0) {
        buffer[19] = '0';
        return buffer[19..20];
    }

    while (num != 0) {
        i -= 1;
        const digit: u8 = '0' + @as(u8, @intCast(num % 10));
        buffer[i] = digit;
        num /= 10;
        std.debug.print("i: {}, digit: {c}\n", .{i, digit});
    }

    const result = buffer[i..20];
    std.debug.print("Result slice: {any}\n", .{result});
    return result;
}