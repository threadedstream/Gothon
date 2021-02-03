from django.urls import path
from apps.api.views import *

urlpatterns = [
    path('save_stats/', SaveStatisticsView.as_view(), name="save_stats"),
    path('retrieve_stats/', RetrieveStatisticsView.as_view(), name="retrieve_stats"),
    path('delete_stats/', DeleteAllStatistics.as_view(), name="delete_stats"),
]

