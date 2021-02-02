from django.db import models
from datetime import date


class Statistics(models.Model):
    date = models.DateField(default=date.today(), null=False)
    views = models.IntegerField(null=True)
    clicks = models.IntegerField(null=True)
    cost = models.FloatField(null=True)