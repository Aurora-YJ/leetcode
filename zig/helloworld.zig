const std = @import("std");

pub fn main() void {
    
    const res: []i32 = ConvertIntegertotheSumofTwoNoZeroIntegers(9);
    for (res) |value| {
        std.debug.print("{}\n", .{value});
    }

}

pub fn ConvertIntegertotheSumofTwoNoZeroIntegers(n: i32) []i32 {
    // var help: i32  = 1;
   for (1..n) |v| {
      const n1: i32 = v; 
      const n2: i32 = n - v;

      if (n1 + n2 == n) {
        std.debug.print("%d %d", .{n1 , n2});
      }
   }
}
