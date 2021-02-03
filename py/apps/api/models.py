from django.db import models
from datetime import date
from django.utils import timezone


class Statistics(models.Model):
    date = models.DateField(null=False)
    views = models.IntegerField(null=True)
    clicks = models.IntegerField(null=True)
    cost = models.FloatField(null=True)
