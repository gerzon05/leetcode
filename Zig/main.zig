const std = @import("std");
const users = @import("user.zig");

const User = users.User;

// String: []const u8
// int
// float
// bool
// void -> vacio

pub fn number_funcion() usize {
    return 5;
}

pub fn sum(value1: i64, value2: i64) i64 {
    return value1 + value2;
}

pub fn main() void {
    const great = "Hola";
    const user1 = User{ .name = "gerzon" };
    const user2 = User{ .name = "sebastian" };

    user1.info();
    user2.info();

    std.debug.print("{s} - {s}\n", .{ great, user1.name });
    std.debug.print("{s}\n", .{"Hello Word"});

    const value_number = number_funcion();
    std.debug.print("{d}\n", .{value_number});

    const value_sum = sum(4, 5);
    std.debug.print("{d}", .{value_sum});
}
