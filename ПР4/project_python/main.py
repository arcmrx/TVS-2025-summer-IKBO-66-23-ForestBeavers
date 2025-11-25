def greet(name):
    print(f"Привет, {name}!")

if __name__ == "__main__":
    # Ошибка 1: NameError
    # undefined_function()

    # Ошибка 2: TypeError
    # "hello" + 5

    # Ошибка 3: IndexError
    my_list = [1, 2, 3]
    # my_list[3]

    # Ошибка 4: AttributeError
    my_string = "Привет"
    my_string.non_existent_attribute

    # Ошибка 5: SyntaxError
    for i in range(5):
        print(i)

    greet("Мир")