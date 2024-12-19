#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

// variables

// char, int, float, double
// Modificadores de acceso:
// const = valor constante -> const int x = 3435;
// volatile = valor es modificado externamente (proceso externo al sistema)

// Modificadores de tipos:
// signed | char - int | con signo por defecto
// unsigned | char - int | sin signo
// log | int - double | largo
// short | int | corto

// Caracteres de modificación

// %c Unico caracter
// %d Valor entero
// %u Entero sin signo en base decimal
// %o Entero en base octal
// %x Entero en base hexadecimal
// %e Un número real en coma flotante, con exponente
// %f Un número real en coma flotante, sin exponente
// %s Valor cadena de caracteres
// %p Dirección de memoria

// #define PI 3.141516 -> otra forma de declarar variables
const float PI = 3.141516;

int main() {
  // Variables
  int number = 12;
  float my_float_number = 12.123;
  bool of_legal_age = false;

  printf("%d\n", number);
  printf("%f\n", my_float_number);
  printf("Hola Mundo en C\n");

  // Pedir valores al usuario
  int age = 0;
  printf("Escribe tu edad: ");
  scanf("%d", &age);
  printf("Tu edad es: %d\n", age);

  // Condicionales
  if (age > 18) {
    printf("Eres mayor de edad\n");
  } else {
    printf("Eres menor de edad\n");
  }

  int genre;
  printf("Selecciona tu genero\n");
  printf("0. Hombre\n");
  printf("1. Mujer\n");
  scanf("%d", &genre);

  switch (genre) {
    case 0:
      printf("Eres Hombre\n");
      break;
    case 1:
      printf("Eres Mujer\n");
      break;
    default:
      printf("No existes en este mundo\n");
      break;
  }

  // Calcular la masa
  float weight, height = 0;

  printf("Escibre tu peso en kg: ");
  scanf("%f", &weight);
  printf("\n");
  printf("Escibre tu altura en metros: ");
  scanf("%f", &height);

  if (weight == 0 || height == 0) {
    printf("No se puede calcular la masa\n");
    return 1;
  };

  float bmi = weight / (height * height);

  if (bmi < 18.5) {
    printf("Se encuentra bajo de peso\n");
  } else if (bmi > 18.5 && bmi < 24.9) {
    printf("Se encuentra en un peso normal\n");
  } else if (bmi > 25 && bmi < 29.9) {
    printf("Se encuentra en un peso muy alto\n");
  } else {
    printf("Se encuentra en obesidad\n");
  };

  // do  While
  int contador = 0;
  do {
    contador++;
  } while (contador < 5);
  printf("%d", contador);
  printf("\n");

  // for
  int *sumanotas = malloc(sizeof(int));
  if (sumanotas == NULL) {
    fprintf(stderr, "Error: espacio de memoria insifucinete.\n");
    return (1);
  }

  *sumanotas = 0;
  for (int i = 0; i < 5; i++) {
    int nota;
    printf("Ingresa la nota: ");
    scanf("%d", &nota);
    *sumanotas += nota;
  };
  printf("La suma de las notas es: %d\n", *sumanotas);

  free(sumanotas);

  // Punteros
  // *: accede al valor de varibale que se esta apuntado
  // &: obtiene la direccion de memomoria de la variable que se esta apuntado
  int valornumero = 10;
  int *valor_puntero = &valornumero;

  // **: apuntar al puntero del puntero
  int **puntero_a_puntero = &valor_puntero;

  // Imprimir la direccion de memoria del puntero
  printf("%p\n", valor_puntero);

  // Imprimir el valor
  printf("%d\n", *valor_puntero);

  // Arrays

  int arrays[3] = {1, 2, 3};

  for (int i = 0; i < 3; i++) {
    printf("%d ", arrays[i]);
  }

  printf("\n");

  return 0;
}
