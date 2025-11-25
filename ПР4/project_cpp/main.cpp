#include <iostream>

void greet(const char* name) {
    std::cout << "Привет, " << name << "!" << std::endl;
}

int divide(int a, int b) {
    if (b == 0) {
        std::cerr << "Ошибка: Деление на ноль!" << std::endl;
        exit(1);
    }
    return a / b;
}

int main() {
    // Ошибка 1: Null pointer dereference
    // greet(nullptr);

    // Ошибка 2: Out-of-bounds access
    int arr[5] = {1, 2, 3, 4, 5};
    // std::cout << arr[5] << std::endl;

    // Ошибка 3: Uninitialized variable
    int uninitialized;
    std::cout << uninitialized << std::endl;


    // Ошибка 4: Division by zero
    divide(1, 0);

    greet("Мир");
    return 0;
}
