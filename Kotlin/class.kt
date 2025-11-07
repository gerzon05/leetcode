class Rectangle(val height: Double, val lenght: Double) {
  val perimeter = (height + lenght) * 2
}

fun main() {
  val value = Rectangle(5.2, 6.7)
  println("EL perimetro es de ${value.perimeter}")
}