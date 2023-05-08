# https://leetcode.com/problems/decoded-string-at-index/
# @param {String} s
# @param {Integer} k
# @return {String}
def decode_at_index(s, k)

  size = 0
  s.chars.each do |c|
    if digit?(c)
      size *= c.to_i
    else
      size += 1
    end
  end

  s.reverse.chars.each do |c|
    k %= size
    if k == 0 and !digit?(c)
      return c
    end


    if digit?(c)
      size /= c.to_i
    else
      size -= 1
    end
  end
end

def digit?(char)
  /\d/ === char
end

s = "leet2code3"              ; k = 10  ; p decode_at_index(s, k)
s = "ha22"                    ; k = 5   ; p decode_at_index(s, k)
s = "a2345678999999999999999" ; k = 1   ; p decode_at_index(s, k)
s = "abc"                     ; k = 1   ; p decode_at_index(s, k)
s = "abc"                     ; k = 3   ; p decode_at_index(s, k)
s = "a23"                     ; k = 6   ; p decode_at_index(s, k)
s = "a2b3c4d5e6f7g8h9"        ; k = 10  ; p decode_at_index(s, k)
s = "a2b3c4d5e6f7g8h9"        ; k = 41  ; p decode_at_index(s, k)
s = "vk6u5xhq9v"              ; k = 554 ; p decode_at_index(s, k)
