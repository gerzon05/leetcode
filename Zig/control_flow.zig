const std = @import("std");
pub fn main() void {
    const Number = 13;

    if (Number >= 18) {
        std.debug.print("Eres Mayor de edad", .{});
    } else {
        std.debug.print("Eres menor de edad", .{});
    }
}
