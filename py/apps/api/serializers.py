from rest_framework import serializers
from .models import Statistics


class StatisticsSerializer(serializers.ModelSerializer):
    date = serializers.DateField(input_formats=['%Y-%m-%d'])
    views = serializers.IntegerField()
    clicks = serializers.IntegerField()
    cost = serializers.FloatField()
    cpc = serializers.CharField(max_length=20, required=False)
    cpm = serializers.CharField(max_length=20, required=False)

    def calculate_cpc_cpm(self, instance) -> None:
        cpc = self.validated_data['cost'] / self.validated_data['clicks']
        cpm = self.validated_data['cost'] / self.validated_data['views'] * 1000

    class Meta:
        model = Statistics
        fields = ('date', 'views', 'clicks', 'cost', 'cpc', 'cpm')


