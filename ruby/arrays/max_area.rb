# @param {Integer[]} height
# @return {Integer}
def max_area(height)
  return 0 if height.empty?
  l = 0
  r = height.size - 1
  max_area = 0
  while l < r
    area = [height[l], height[r]].min * (r - l)
    p [l,r,area,max_area]
    if area > max_area
      max_area = area
    end
    if height[l] <= height[r]
      l += 1
    else
      r -= 1
    end
  end

  max_area
end

height = [1,8,6,2,5,4,8,3,7]
p max_area(height)
