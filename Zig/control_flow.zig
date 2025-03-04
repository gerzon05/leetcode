const std = @import("std");
pub fn main() void {
    const Number = 13;

    // const arrayNumber = [10]u8{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }; No puede cambiar su valor
    var arrayNumber = [10]u8{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };

    const newArray = arrayNumber[0..arrayNumber.len]; // no importa el tipo depende del tipo de arrays anterior

    newArray[2] = 34;

    var name = "Gerzon";
    name = "Rangel";

    std.debug.print("{d}\n", .{newArray[2]});
    std.debug.print("{s}\n", .{name});

    if (Number >= 18) {
        std.debug.print("Eres Mayor de edad\n", .{});
    } else {
        std.debug.print("Eres menor de edad\n", .{});
    }

    const method = "GET";

    if (std.mem.eql(u8, method, "GET") or std.mem.eql(u8, method, "HEAD")) {
        std.debug.print("El metodo para obtener datos es: {s}\n", .{method});
    } else if (std.mem.eql(u8, method, "POST")) {
        std.debug.print("El metodo para guardar datos es: {s}\n", .{method});
    } else {
        std.debug.print("No eligiste un metodo valido\n", .{});
    }

    const rol = "Admin";

    if (std.ascii.eqlIgnoreCase(rol, "Usuario")) {
        std.debug.print("Estas logueado con el rol Usuario\n", .{});
    } else if (std.ascii.eqlIgnoreCase(rol, "Admin")) {
        std.debug.print("Estas logueado con el rol Admin\n", .{});
    } else {
        std.debug.print("No existe ese Rol\n", .{});
    }

    const numbers = [5]usize{ 45, 65, 23, 87, 12 };

    for (0..numbers.len) |value| {
        std.debug.print("{d} ", .{numbers[value]});
    }
    std.debug.print("\n", .{});

    const number = 1;

    switch (number) {
        1 => return std.debug.print("paper", .{}),
        2 => return std.debug.print("cotton", .{}),
        3 => return std.debug.print("leather", .{}),
        4 => return std.debug.print("flower", .{}),
        5 => return std.debug.print("wood", .{}),
        6 => return std.debug.print("sugar", .{}),
        else => return std.debug.print("Does not exist", .{}),
    }
}
