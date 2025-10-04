const std = @import("std");

pub fn main() void {
    const  allocator = std.heap.page_allocator;    

    var  mymap1 = std.AutoHashMap([]const u8 , std.ArrayList([]const u8)).init(allocator);

    var first1 = std.ArrayList([]const u8).init(allocator);
    try first1.append("Sakura");
    try first1.append("Hinata");

    try mymap1.put("Naruto", first1);
    try mymap1.put("Sasuke", first1);

//    std.debug.println("{} -> ",  .{mymap1});

    // var it = mymap1.iterator();
    // while (it.next()) |entry| {
    //     std.debug.print("{} -> ", .{entry.key});
    //     var list_it = entry.value.iterator();
    //     while (list_it.next()) |friend| {
    //         std.debug.print("{} ", .{friend});
    //     }
    //     std.debug.print("\n", .{});
    // }


} 