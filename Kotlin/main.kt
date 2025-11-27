fun main() {
    println("Hello, World!")

    // ======= VARIABLES ================
    // var = variable que puede cambiar
    // val = variable inmutable (constante)

    var number: Int = 5
    number += 1

    println("El valor de la variable es de: "+ number)

    val value: Int

    value = 4

    println("El valor de la variable asignada: "+ value)


    // =========== STRING TEMPLATES ================

    var age: Int = 20
    val sentence_one = "tu edad es de $age"
    age = 21
    val sentence_two = "${sentence_one.replace("tu", "Tu")}, y cumpliras $age años"

    println(sentence_two)


    // ======= CODICIONALES ============

    println("EL numero mayor entre 10 y 40 es: ${maxOf(10,40)}")

    println("\n ============== FOR LOPP ============== \n")

    // ======= ARRAYS ================
    val itemsFruits = listOf("apple", "banana", "kiwifruit")

    // ====== FOR LOPP ==============
    for (itemFruit in itemsFruits){
        println(itemFruit)
    }

    println("\n ============== FOR LOPP INDEX ============== \n")

    // ====== FOR LOPP INDEX ==============
    for (index in itemsFruits.indices){
        println("La posicion $index es la fruta ${itemsFruits[index]}")
    }
}

// ======= CODICIONALES ============
fun maxOf(a: Int, b: Int): Int {
    if(a > b){
        return a
    }else {
        return b
    }
}

// fun maxOf(a: Int, b: Int) = if (a > b) a else b -> forma corta


// crear el archivo ejecutables
// kotlinc main.kt -include-runtime -d main.jar
// java -jar main.jar 