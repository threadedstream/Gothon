def sumToFloat(sum: str) -> float:
    parts = sum.split(' ')
    rubs = parts[0][0:len(parts[0]) - 1].strip()
    if len(parts) > 1:
        kops = parts[1][0:len(parts[1]) - 1].strip()
    else:
        kops = "0"

    return float(rubs + "." + kops)
