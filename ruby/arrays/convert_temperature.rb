# https://leetcode.com/problems/convert-the-temperature/
# Kelvin = Celsius + 273.15
# Fahrenheit = Celsius * 1.80 + 32.00

# @param {Float} celsius
# @return {Float[]}
def convert_temperature(celsius)
  kelvin = celsius + 273.15
  fahrenheit = celsius * 1.80 + 32.00
  [kelvin, fahrenheit]
end

p convert_temperature(36.5)
