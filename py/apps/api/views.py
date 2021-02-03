from django.db.models import Q
from rest_framework import status
from rest_framework.generics import CreateAPIView, ListAPIView, DestroyAPIView
from rest_framework.permissions import AllowAny
from rest_framework.response import Response
from .models import Statistics
import re

from rest_framework.views import APIView

from .utils import cost_to_float, alter_json
from .serializers import StatisticsSerializer


class SaveStatisticsView(CreateAPIView):
    serializer_class = StatisticsSerializer
    permission_classes = (AllowAny,)

    def post(self, request, *args, **kwargs):
        cost = 0
        # Avoid division by zero
        views = 1
        clicks = 1
        data = {}
        if 'date' not in request.data:
            data['success'] = False
            data['result'] = "Missing mandatory field date"
            return Response(data=data, status=status.HTTP_400_BAD_REQUEST)
        else:
            date = request.data['date']
        if 'views' in request.data:
            views = int(request.data['views'])
        if 'clicks' in request.data:
            clicks = int(request.data['clicks'])
        if 'cost' in request.data:
            cost = cost_to_float(request.data['cost'])

        s_data = {'date': date, 'views': views, 'clicks': clicks, 'cost': cost}
        s = StatisticsSerializer(data=s_data)
        if s.is_valid():
            model = s.save()
            data['success'] = True
            data['result'] = "OK"
            return Response(data=data, status=status.HTTP_201_CREATED)
        else:
            data['success'] = False
            data['result'] = s.errors
            return Response(data=data, status=status.HTTP_400_BAD_REQUEST)


class RetrieveStatisticsView(APIView):
    serializer_class = StatisticsSerializer
    permission_classes = (AllowAny,)

    def get(self, request, *args, **kwargs):
        order_by = ''
        data = {}
        if 'to' not in request.GET and 'from' not in request.GET:
            data['success'] = False
            data['result'] = "Please, specify all necessary fields: 'to' and 'from'"
            return Response(data=data, status=status.HTTP_400_BAD_REQUEST)

        to = request.GET['to']
        fr = request.GET['from']
        if 'order_by' in request.GET:
            order_by = request.GET['order_by']
        if order_by != 'date' and order_by != 'views' and order_by != 'clicks' and order_by != 'cost':
            # Default option
            order_by = 'date'

        statistics = StatisticsSerializer(
            Statistics.objects.filter(Q(date__lte=to), Q(date__gte=fr)).order_by(order_by), many=True).data
        data['success'] = True
        data['result'] = alter_json(statistics)
        return Response(data=data, status=status.HTTP_200_OK)


class DeleteAllStatistics(DestroyAPIView):
    serializer_class = StatisticsSerializer
    permission_classes = (AllowAny,)

    def delete(self, request, *args, **kwargs):
        data = {}
        Statistics.objects.all().delete()
        data['success'] = True
        data['result'] = "OK"
        return Response(data=data, status=status.HTTP_200_OK)
