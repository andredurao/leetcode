# @param {String} s
# @return {String}
def longest_palindrome(s)
  cursors = [0,0]
  dist = 0

  0.upto(s.size) do |shift|
    odd_length = expand(shift, shift, s)
    if odd_length > cursors[1] - cursors[0] + 1
      dist = odd_length / 2
      cursors = [shift - dist, shift + dist]
    end

    even_length = expand(shift, shift + 1, s)
    if even_length > cursors[1] - cursors[0] + 1
      dist = (even_length / 2) - 1
      cursors = [shift - dist, shift + 1 + dist]
    end
  end

  i, j = cursors
  s[i..j]
end

def expand(i, j, s)
  l = i ; r = j

  while l >= 0 && r < s.size && s[l] == s[r]
    l -= 1
    r += 1
  end
  r - l - 1
end

p longest_palindrome("babad")
p longest_palindrome("cbbd")
p longest_palindrome("a")

s = "pihoigwlvzvtrugdolvtzrkyelaqdvbijzmkhebzawboaxkdjyfocpewwztffuaibcqurwwmijmvrnpfcoglyxpxkrbhupoxcafabxtoecodsjgngrionuvzaiigevuvruxxiwpjzjlqgenglhprcgzgpdzabrjhkbtfrbmwpcszepxhwiwdhvnpsmhhaiqsbeiwsaeomqtzcpjzfknejxlxwtpkufanhuoyjgihdzhtxnyctazzvnttjspfztjurdwmmzrvobcatkorfhpieoqfetcglembkgbexsznuduhrfoxkbswkanqwfkoktnnujqetijaqhrxuhkgsezfdrncbaltctwcourdbpdwhqlsxfwsoaduaqkbjeekwwykptjthhtokrvzsuelsywyznqscnwiszogzqvfsgggzltlvzkllinpfaigswquqfvabbzvestwxhbnfjhnvfhyxalchmtkcwnyyrbwjsoqooypwteozbivqiyldpqlykxinmzkgnmfbobgjivlzubfen"
p longest_palindrome(s)
