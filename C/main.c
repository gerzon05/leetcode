#include <stdbool.h>
#include <stdio.h>

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

#define PI 3.141516

int main(){
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
  if (age > 18){
   printf("Eres mayor de edad\n");
  } else {
    printf("Eres menor de edad\n");
  }

  int genre;
  printf("Selecciona tu genero\n");
  printf("0. Hombre\n");
  printf("1. Mujer\n");
  scanf("%d", &genre);

  switch (genre){
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

  if (weight == 0 || height == 0){
    printf("No se puede calcular la masa\n");
    return 1;
  };

  float bmi = weight / (height * height);

  if (bmi < 18.5){
    printf("Se encuentra bajo de peso\n");
  } else if (bmi > 18.5 && bmi < 24.9){
    printf("Se encuentra en un peso normal\n");
  }else if (bmi > 25 && bmi < 29.9){
    printf("Se encuentra en un peso muy alto\n");
  } else {
    printf("Se encuentra en obesidad\n");
  };

  //do  While
  int contador = 0;
  do
  {
    contador++;
  } while (contador < 5);
  printf("%d", contador);

  // for
  int nota, sumanotas = 0;
  for (int i = 0; i < 5; i++)
  {
    printf("Ingresa la nota: ");
    scanf("%d", &nota);
    sumanotas += nota;
  };
  printf("%d", sumanotas)

  return 0;
}


