# @param {Integer[][]} img1
# @param {Integer[][]} img2
# @return {Integer}
def largest_overlap(img1, img2)
  bin_img1 = convert_to_int(img1)
  bin_img2 = convert_to_int(img2)
  size = img1[0].size

  # binding.irb
  # exit 0
  max = i = 0
  while i < size
    j = 0
    while j < size
      [:d, :u].each do |vertical|
        [:l, :r].each do |horizontal|
          img_lookup = bin_img1.dup
          img_lookup = shift(img_lookup, size, horizontal, j)
          img_lookup = shift(img_lookup, size, vertical, i)
          max = [
            max,
            diff_counter(img_lookup, bin_img2, size),
          ].max
        end
      end
      j += 1
    end
    i += 1
  end
  max
end

# move img left, right, up & down (only L & R required here)
def shift(img, size, direction, units)
  return img if units == 0
  result = []
  img.each_with_index do |row, index|
    if direction == :l
      result << (row << units)
    elsif direction == :r
      result << (row >> units)
    elsif direction == :u
      shift = index + units # [0 1 2] ~> [1 2 x]
      result << (img[shift] || 0)
    elsif direction == :d
      shift = index - units # [0 1 2] ~> [x 0 1]
      result << (shift < 0 ? 0 : img[shift])
    end
  end
  result
end

def convert_to_int(img)
  img.reduce([]){|array, row| array << row.join.to_i(2) ; array}
end

def print_img(img)
  p img.map{|row| row.to_s(2)}
end

# :.| unfortunately, Ruby doesn't have a popcount function like other languages
# Brian Kernighan’s popcount
def popcount(n)
  count = 0
  while n > 0
    n &= (n - 1)
    count += 1
  end
  count
end

def diff_counter(bin_img1, bin_img2, size)
  total = 0
  bin_img1.each_with_index do |val1, index|
    total += popcount(val1 & bin_img2[index])
  end
  total
end

# popcount with remainders of 2*4 using a pre-computed hash
def popcount_using_hash_nibbles(n)
end

# TODO 1 try out with the pre-computed hash to check if there's any improvement
# TODO 2 instead of bitcounting the diffs on each time, cache the diff on every call using a hash of hashes, eg:
#        diff(1, 3) => hash[1] => { 3 => (1 the value) } and hash[3] => { 1 => (1 the value) }


img1 = [ [1,1,0],
         [0,1,0],
         [0,1,0] ]
img2 = [ [0,0,0],
         [0,1,1],
         [0,0,1] ]
result = largest_overlap(img1, img2)
p result


img1 = [[1]]
img2 = [[1]]
result = largest_overlap(img1, img2)
p result


img1 = [[0]]
img2 = [[0]]
result = largest_overlap(img1, img2)
p result

img1 = [[0]]
img2 = [[1]]
result = largest_overlap(img1, img2)
p result


img1 = [[0,0,0,0,1],
        [0,0,0,0,0],
        [0,0,0,0,0],
        [0,0,0,0,0],
        [0,0,0,0,0]]
img2 = [[0,0,0,0,0],
        [0,0,0,0,0],
        [0,0,0,0,0],
        [0,0,0,0,0],
        [1,0,0,0,0]]
result = largest_overlap(img1, img2)
p result
