# @param {Integer} target
# @return {Integer}
def racecar(target)
  a_commands = []
  move_helper(target, a_commands, 0, 1, 'A', 0)
  r_commands = []
  move_helper(target, r_commands, 0, 1, 'R', 0)
  p a_commands
  p r_commands
end

def move_helper(target, commands=[], pos=0, speed=1, sequence='', previous_pos=0)
  p [target, pos, speed, sequence, previous_pos] ; sleep 0.3
  new_pos = pos
  if pos == target
    commands << sequence
    return
  else
    direction = (target - previous_pos) - (target - pos)
    if direction < 0
      p 'wrong way'
      return
    else
      p 'right way'
      if target >= 0
        return if pos > target
      else
        return if pos < target
      end
    end
  end
  command = sequence[-1]
  if command == 'A'
    new_pos += speed
    speed *= 2
  elsif command == 'R'
    speed = speed > 0 ? -1 : 1
  end
  move_helper(target, commands, new_pos, speed, sequence+'A', pos)
  move_helper(target, commands, new_pos, speed, sequence+'R', pos)
end

target = 3
p racecar(target)
