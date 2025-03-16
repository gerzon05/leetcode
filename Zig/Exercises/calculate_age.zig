// Objetivo: Crear una función que calcule la edad usando variables y operaciones básicas.

// Input: birth_year = 1990, current_year = 2024
// Output esperado: 34
const std = @import("std");

fn calculate_age(birthYear: u64, currentYear: u64) !u64 {
    return currentYear - birthYear;
}

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    const stdin = std.io.getStdIn().reader();

    var buffer: [100]u8 = undefined;

    try stdout.print("Year of birth: ", .{});
    const birthYearString = try stdin.readUntilDelimiter(&buffer, '\n');

    try stdout.print("Current year: ", .{});
    const currentYearString = try stdin.readUntilDelimiter(&buffer, '\n');

    const birthYear = try std.fmt.parseInt(u64, birthYearString, 10);
    const currentYear = try std.fmt.parseInt(u64, currentYearString, 10);

    const result = calculate_age(birthYear, currentYear);
    try stdout.print("You are {d} years old\n", .{result});
}
