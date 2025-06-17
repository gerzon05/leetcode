#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int error = 0;

#define OK 0
#define ERROR_EXCEEDED_LIMIT 1
#define ERROR_INVALID_NUMBER 2
#define ERROR_INVALID_INPUT 3
#define ERROR_MEMORY_ALLOCATION 4

const char* message(int code) {
  switch (code) {
    case OK:
      return "No tiene error";
    case ERROR_EXCEEDED_LIMIT:
      return "El número ingresado no debe ser mayor a 10";
    case ERROR_INVALID_NUMBER:
      return "El número ingresado debe ser mayor a 0";
    case ERROR_INVALID_INPUT:
      return "El valor ingresado es inválido";
    case ERROR_MEMORY_ALLOCATION:
      return "Error de memoria";
    default:
      return "Error desconocido";
  }
}

typedef struct Students {
  char name[50];
  float ratings[3];
  float average;
} Student;

int count_number_of_students() {
  char buffer[100];

  printf("Ingrese la cantidad de estudiantes: ");

  if (!fgets(buffer, sizeof(buffer), stdin)) {
    error = ERROR_INVALID_INPUT;
    return 0;
  }

  buffer[strcspn(buffer, "\n")] = '\0';

  if (buffer[0] == '\0') {
    error = ERROR_INVALID_INPUT;
    return 0;
  }

  char* conversion_finished;
  long number_of_students = strtol(buffer, &conversion_finished, 10);

  if (*conversion_finished != '\0' || conversion_finished == buffer) {
    error = ERROR_INVALID_INPUT;
    return 0;
  }

  if (number_of_students <= 0) {
    error = ERROR_INVALID_NUMBER;
    return 0;
  }

  if (number_of_students > 10) {
    error = ERROR_EXCEEDED_LIMIT;
    return 0;
  }

  error = OK;
  return number_of_students;
}

float get_student_rating(const char* message_prompt) {
  char buffer[100];
  double number;

  printf("%s", message_prompt);

  if (fgets(buffer, sizeof(buffer), stdin) == NULL) {
    error = ERROR_INVALID_INPUT;
    return 0;
  }

  buffer[strcspn(buffer, "\n")] = '\0';

  if (buffer[0] == '\0') {
    error = ERROR_INVALID_INPUT;
    return 0;
  }

  char* conversion_finished;
  number = strtod(buffer, &conversion_finished);

  if (*conversion_finished != '\0' || conversion_finished == buffer) {
    error = ERROR_INVALID_INPUT;
    return 0;
  }

  if (number <= 0) {
    error = ERROR_INVALID_NUMBER;
    return 0;
  }

  if (number > 10) {
    error = ERROR_EXCEEDED_LIMIT;
    return 0;
  }

  error = OK;
  return number;
}

Student* get_students_with_average(int amount_students) {
  Student* students = NULL;
  students = (Student*)malloc(amount_students * sizeof(Student));

  if (students == NULL) {
    error = ERROR_MEMORY_ALLOCATION;
    return NULL;
  }

  for (int i = 0; i < amount_students; i++) {
    printf("Ingrese el nombre del estudiante %d: ", i + 1);

    double average = 0;

    if (!fgets(students[i].name, sizeof(students[i].name), stdin)) {
      error = ERROR_INVALID_INPUT;
      free(students);
      return NULL;
    }

    students[i].name[strcspn(students[i].name, "\n")] = 0;

    for (int r = 0; r < 3; r++) {
      char message_prompt[100];
      sprintf(message_prompt, "Ingrese la nota %d del estudiante %s: ", r + 1,
              students[i].name);
      float rating = get_student_rating(message_prompt);
      if (error != OK) {
        printf("Error: %s\n", message(error));
        free(students);
        return NULL;
      }
      students[i].ratings[r] = rating;
      average += rating;
    }

    students[i].average = average / 3;
  }

  return students;
}

Student find_student_with_max_average(Student* students, int total_students) {
  Student student_with_max_average = students[0];
  int index_of_best_student = 0;

  for (int i = 1; i < total_students; i++) {
    if (students[index_of_best_student].average < students[i].average) {
      index_of_best_student = i;
      student_with_max_average = students[i];
    }
  }

  return student_with_max_average;
}

float calculate_class_average(Student* students, int total_students) {
  float total_class_average = 0;

  for (int i = 0; i < total_students; i++) {
    total_class_average += students[i].average;
  }

  return total_class_average / total_students;
}

int main() {
  int total_students = count_number_of_students();

  if (error != OK) {
    printf("Error: %s\n", message(error));
    return 0;
  }

  printf("\n");
  Student* students = get_students_with_average(total_students);

  if (students == NULL) {
    printf("Error: %s\n", message(error));
    return 0;
  }

  for (int i = 0; i < total_students; i++) {
    printf("\n");
    printf("Nombre del Estudiante: %s\n", students[i].name);
    printf("Calificaciones: %.1f, %.1f, %.1f\n", students[i].ratings[0],
           students[i].ratings[1], students[i].ratings[2]);
    printf("Promedio: %.2f\n", students[i].average);
    printf("----------------------------\n");
  }

  Student student_with_best_average =
      find_student_with_max_average(students, total_students);

  printf("\n");
  printf("Estudiante con el mejor promedio de la clase: \n");
  printf("Nombre del Estudiante: %s\n", student_with_best_average.name);
  printf("Calificaciones: %.1f, %.1f, %.1f\n",
         student_with_best_average.ratings[0],
         student_with_best_average.ratings[1],
         student_with_best_average.ratings[2]);
  printf("Promedio: %.2f\n", student_with_best_average.average);
  printf("----------------------------\n");

  float class_average_result =
      calculate_class_average(students, total_students);

  printf("\n");
  printf("El promedio de la clase es de: %.2f\n", class_average_result);
  printf("----------------------------\n");

  free(students);

  return 0;
}