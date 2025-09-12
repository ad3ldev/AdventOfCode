const std = @import("std");

pub fn main() !void {
    const filename = "input.txt";

    // Create allocator
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    // Read file
    const file = std.fs.cwd().openFile(filename, .{}) catch |err| {
        std.debug.print("Error opening file: {}\n", .{err});
        return;
    };
    defer file.close();

    const content = file.readToEndAlloc(allocator, std.math.maxInt(usize)) catch |err| {
        std.debug.print("Error reading file: {}\n", .{err});
        return;
    };
    defer allocator.free(content);

    // Count parentheses
    var count: i32 = 0;
    for (content) |char| {
        switch (char) {
            '(' => count += 1,
            ')' => count -= 1,
            else => {},
        }
    }

    std.debug.print("{}\n", .{count});
}
