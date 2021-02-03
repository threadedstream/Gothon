from django.test import TestCase
from .utils import *
from rest_framework.test import APIClient, APITestCase


class UtilsTestCase(TestCase):
    def setUp(self) -> None:
        self.sample1 = 12.1
        self.sample2 = 12.34
        self.sample3 = 0.32
        self.sample4 = 12.78
        self.num1 = '12r 1k'
        self.num2 = '12r 34k'
        self.num3 = '0r 32k'
        self.num4 = '12r 78k'

    def test_cost_to_float(self):
        number1 = cost_to_float(self.num1)
        number2 = cost_to_float(self.num2)
        number3 = cost_to_float(self.num3)
        number4 = cost_to_float(self.num4)
        self.assertEqual(number1, 12.1)
        self.assertEqual(number2, 12.34)
        self.assertEqual(number3, 0.32)
        self.assertEqual(number4, 12.78)

    def test_float_to_cost(self):
        res1 = float_to_cost(self.sample1)
        res2 = float_to_cost(self.sample2)
        res3 = float_to_cost(self.sample3)
        res4 = float_to_cost(self.sample4)
        self.assertEqual(res1, '12r 1k')
        self.assertEqual(res2, '12r 34k')
        self.assertEqual(res3, '0r 32k')
        self.assertEqual(res4, '12r 78k')


class SaveStatisticsRequestTestCase(APITestCase):
    def setUp(self) -> None:
        self.url = "http://192.168.1.41:8000/api/v1/save_stats/"
        self.date = '2017-12-20'
        self.clicks = 100
        self.views = 240
        self.cost = "12r 20k"

    def test_save_statistics(self):
        client = APIClient()
        data = {
            'date': self.date,
            'clicks': self.clicks,
            'views': self.views,
            'cost': self.cost
        }
        response = client.post(self.url, data)
        self.assertEqual(response.status_code, 201)


class RetrieveStatisticsRequestTestCase(APITestCase):
    def setUp(self) -> None:
        self.url = "http://192.168.1.41:8000/api/v1/retrieve_stats/"
        self.to = "2078-12-01"
        self.fr = "2077-12-30"
        self.order_by = "clicks"

    def test_retrieve_statistics(self):
        client = APIClient()
        data = {
            'to': self.to,
            'from': self.fr,
            'order_by': self.order_by
        }

        response = client.get(self.url, data)
        self.assertEqual(response.status_code, 200)


class DeleteAllStatisticsRequestTestCase(APITestCase):
    def setUp(self) -> None:
        self.url = "http://192.168.1.41:8000/api/v1/delete_stats/"

    def test_delete_statistics(self):
        client = APIClient()
        response = client.delete(self.url)
        self.assertEqual(response.status_code, 200)
