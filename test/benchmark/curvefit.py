from typing import List
from scipy.optimize import curve_fit

def func(x, a, b):
    return a*x + b

def getPopt(trainX: List[int], trainY: List[int]) -> List[float]:
    popt, _ = curve_fit(func, trainX, trainY)
    return [float(x) for x in popt]