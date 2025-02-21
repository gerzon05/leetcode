#include <math.h>>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Ejercicios para Tareas

// Libro: Define una estructura llamada Libro que contenga los siguientes
// miembros: titulo (cadena), autor (cadena), anio (entero) y precio (flotante).
// Crea un programa que declare una variable de tipo Libro, asigne valores a sus
// miembros y luego imprima la información del libro.

// Rectángulo: Define una estructura llamada Rectangulo que contenga los
// siguientes miembros: ancho (flotante) y alto (flotante). Crea una función que
// reciba un Rectangulo como argumento y calcule su área. Crea un programa que
// declare una variable de tipo Rectangulo, asigne valores a sus miembros, llame
// a la función para calcular el área e imprima el resultado.

// Estudiante: Define una estructura llamada Estudiante que contenga los
// siguientes miembros: nombre (cadena), materia (cadena) y nota
// (int). Crea un array de 5 estudiantes. Pide al usuario que ingrese la
// información de cada estudiante y almacénala en el array. Luego, calcula e
// imprime el promedio general de todos los estudiantes.

// Coordenada: Define una estructura llamada Coordenada con miembros x e y
// (ambos enteros). Crea una función que reciba dos estructuras Coordenada y
// calcule la distancia entre ellas. La distancia se calcula como la raíz
// cuadrada de (x2 - x1)^2 + (y2 - y1)^2. Puedes usar la función sqrt() de la
// librería math.h.

struct Product {
  char title[50];
  int price;
};

struct Persona {
  char name[50];
  int age;

  struct Product shopping_cart;
};

struct Plan {
  struct Persona user;
  struct Product product;
};

typedef struct Book {
  char title[50];
  char author[50];
  int year;
  float price;
} Book;

typedef struct Rectangle {
  float width;
  float height;
} Rectangle;

typedef struct Student {
  char name[50];
  char subject[50];
  int note;
} Student;

typedef struct Coordiante {
  int x;
  int y;
} Coordiante;

float calculateArea(Rectangle data) { return data.width * data.height; };

int calulateCordinate(Coordiante value1, Coordiante value2) {
  int resCoordinate1 = value2.x - value1.x;
  int resCoordinate2 = value2.y - value1.y;
  int coorinate1 = sqrt(resCoordinate1);
  int coorinate2 = sqrt(resCoordinate2);
  return coorinate1 + coorinate2;
};

int main() {
  struct Persona Persona1;
  struct Plan plan;
  struct Persona *persona2 = &Persona1;

  Book book;

  strcpy(book.title, "ADN BARCA");
  strcpy(book.author, "Paco Seirullo");
  book.year = 2024;
  book.price = 109.999;

  printf(
      "------------------------------------  Libro  "
      "------------------------------------\n");
  printf("Title: %s\n", book.title);
  printf("Title: %s\n", book.author);
  printf("Title: %d\n", book.year);
  printf("Title: %u\n", book.price);
  printf(
      "------------------------------------------------------------------------"
      "---------"
      "\n");

  Rectangle rectangle;

  rectangle.width = 2;
  rectangle.height = 4;

  float resultArea = calculateArea(rectangle);

  printf(
      "------------------------------------  Rectangulo  "
      "------------------------------------\n");

  printf("El area del rectangulo es: %f\n", resultArea);

  printf(
      "------------------------------------------------------------------------"
      "---------"
      "\n");

  const int LENT_STUDENT = 1;

  char subject[50];
  Student *studens = (Student *)malloc(LENT_STUDENT * sizeof(char));

  if (studens == NULL) {
    fprintf(stderr, "Error: espacio de memoria insifucinete.\n");
    return (1);
  }

  printf("Escibre la materia: \n");
  scanf("%s", &subject);

  for (int i = 0; i < LENT_STUDENT; i++) {
    char name[50];
    int note;
    printf("\nEscibre el nombre del estudiante: ");
    scanf("%s", &studens[i].name);
    printf("Escibre la nota: ");
    scanf("%d", &studens[i].note);
    strcpy(studens[i].subject, subject);
  };

  for (int i = 0; i < LENT_STUDENT; i++) {
    int number = i + 1;
    printf("---------- Estudiante %d -----------------\n", number);
    printf("nombre del estudiante: %s\n", studens[i].name);
    printf("Materia: %s\n", studens[i].subject);
    printf("nota del estudiante: %d\n", studens[i].note);
    printf("-----------------------------------------\n");
  };

  free(studens);

  Coordiante coordiante1;
  Coordiante coordiante2;

  coordiante1.x = 5;
  coordiante1.y = 10;

  coordiante2.x = 6;
  coordiante2.y = 10;

  int resultCoordinate = calulateCordinate(coordiante1, coordiante2);

  printf("---------- Calcular Cordenada -----------------\n");
  printf("la distancia entre las coordenadas es de: %d\n", resultCoordinate);
  printf("-----------------------------------------\n");

  strcpy(Persona1.name, "Gerzon Rangel");
  strcpy(Persona1.shopping_cart.title, "Nissan 19");
  Persona1.age = 19;

  plan.product.price = 10000;

  printf("Nombre de (referencia) %s \n", persona2->name);
  printf("%s \n", Persona1.name);
  printf("%d \n", Persona1.age);

  return (0);
}
