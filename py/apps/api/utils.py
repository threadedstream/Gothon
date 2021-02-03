from typing import Union, Dict

from rest_framework.utils.serializer_helpers import ReturnDict


def cost_to_float(cost: str) -> float:
    parts = cost.split(' ')
    rubs = parts[0][0:len(parts[0]) - 1].strip()
    if len(parts) > 1:
        kops = parts[1][0:len(parts[1]) - 1].strip()
    else:
        kops = "0"

    return float(rubs + "." + kops)


def float_to_cost(num: float) -> str:
    numStr = str(num)
    parts = numStr.split('.')
    rubs = parts[0] + "r"
    if len(parts) > 1:
        kops = parts[1] + "k"
    else:
        kops = "0k"

    return rubs + " " + kops


def alter_json(data: ReturnDict) -> ReturnDict:
    for d in data:
        d['cpc'] = float_to_cost(d['cost'] / d['clicks'])
        d['cpm'] = float_to_cost(d['cost'] / d['views'] * 1000)
        d['cost'] = float_to_cost(d['cost'])
    return data
