const std = @import("std");
pub fn main() void {
    const Number = 13;

    // const arrayNumber = [10]u8{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }; No puede cambiar su valor
    var arrayNumber = [10]u8{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };

    const newArray = arrayNumber[0..arrayNumber.len]; // no importa el tipo depende del tipo de arrays anterior

    newArray[2] = 34;

    std.debug.print("{d}\n", .{newArray[2]});

    if (Number >= 18) {
        std.debug.print("Eres Mayor de edad\n", .{});
    } else {
        std.debug.print("Eres menor de edad\n", .{});
    }
}
