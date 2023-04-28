# @param {Integer[]} ratings
# @return {Integer}
def candy(ratings)
  # start with one for each rating
  l2r = [1] * ratings.size
  r2l = [1] * ratings.size
  last = ratings.size - 1

  1.upto(last) do |i|
    # L2R
    if ratings[i] > ratings[i - 1]
      l2r[i] = l2r[i - 1] + 1
    end

    # R2L
    if ratings[last - i] > ratings[last - i + 1]
      r2l[last - i] = r2l[last - i + 1] + 1
    end
  end

  candies = []
  l2r.each_index{|i| candies << [l2r[i],r2l[i]].max}

  candies.sum
end

def neighbours_indexes(ratings, i)
  [i - 1, i + 1].delete_if{|x| x < 0 || x >= ratings.size}
end

def fetch(array, i)
  return nil if !i || i < 0 || i >= array.size
  array[i]
end

ratings = [1,0,2]
p candy(ratings)

ratings = [1,2,2]
p candy(ratings)

ratings = [0,1,2,3,4,4,4]
p candy(ratings)

ratings = [29,51,87,87,72,12]
p candy(ratings)

