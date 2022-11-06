# @param {String} ransom_note
# @param {String} magazine
# @return {Boolean}
def can_construct(ransom_note, magazine)
  magazine_map = magazine.chars.tally
  ransom_note.chars.tally.each do |char, count|
    if !magazine_map[char] || magazine_map[char] < count
      return false
    end
  end
  true
end
