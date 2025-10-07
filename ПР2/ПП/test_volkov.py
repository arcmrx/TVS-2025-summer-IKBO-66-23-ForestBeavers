import unittest
import volkov


class TestVolkov(unittest.TestCase):
    def setUp(self):
        self.data = [5, 3, 8, 1, 9]

    def test_find_max(self):
        self.assertEqual(volkov.find_max(self.data), 9, "Максимальное значение должно быть 9")

    def test_find_min(self):
        self.assertEqual(volkov.find_min(self.data), 1, "Минимальное значение должно быть 1")

    def test_sort_array(self):
        self.assertEqual(volkov.sort_array(self.data), [1, 3, 5, 8, 9], "Массив должен быть отсортирован по возрастанию")

    def test_reverse_array(self):
        self.assertEqual(volkov.reverse_array(self.data), [9, 1, 8, 3, 5], "Массив должен быть перевернут")

    def test_get_average(self):
        # Ожидаем среднее значение (а не сумму)
        expected = (sum(self.data) / len(self.data))
        result = volkov.get_average(self.data)
        self.assertEqual(result, expected, "Ошибка: функция возвращает сумму вместо среднего")


if __name__ == '__main__':
    unittest.main()
