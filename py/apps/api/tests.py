from django.test import TestCase
from .utils import *

class UtilsTestCase(TestCase):
    def setUp(self) -> None:
        self.num1 = '12.1'
        self.num2 = '12.134'
        self.num3 = '.132'
        self.num4 = '0.0.0'

    def test_sum_to_float(self):
        number1 = sumToFloat(self.num1)
        number2 = sumToFloat(self.num2)
        number3 = sumToFloat(self.num3)
        number4 = sumToFloat(self.num4)
        self.assertEqual(number1, 12.1)
        self.assertEqual(number2, 12.134)
        self.assertEqual(number3, 0.132)
        self.assertEqual(number4, None)