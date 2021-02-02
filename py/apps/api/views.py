from rest_framework import status
from rest_framework.generics import CreateAPIView
from rest_framework.response import Response


class SaveStatisticsView(CreateAPIView):
    def post(self, request, *args, **kwargs):
        data = {}
        if 'date' not in request.data:
            data['success'] = False
            data['result'] = "date field must be present"
            return Response(data=data, status=status.HTTP_400_BAD_REQUEST)
        if 'views' in request.data:
            views = int(request.data['views'])
        if 'clicks' in request.data:
            clicks = int(request.data['clicks'])
        if 'cost' in request.data:

