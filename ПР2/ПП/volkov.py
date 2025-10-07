def find_max(arr):
    if not arr:
        raise ValueError("Массив пуст")
    return max(arr)


def find_min(arr):
    if not arr:
        raise ValueError("Массив пуст")
    return min(arr)


def sort_array(arr):
    return sorted(arr)


def reverse_array(arr):
    return arr[::-1]


def get_average(arr):
    if not arr:
        raise ValueError("Массив пуст")
    # Ошибка: должно быть sum(arr) / len(arr)
    return sum(arr)