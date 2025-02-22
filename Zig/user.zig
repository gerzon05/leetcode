const std = @import("std");

pub const User = struct {
    name: []const u8,

    pub fn info(self: User) void {
        std.debug.print("{s}\n", .{self.name});
    }
};
