# @param {Integer[][]} rooms
# @return {Boolean}
def can_visit_all_rooms(rooms)
  keys = []
  map = {}
  rooms.each_with_index{|keys, room| map[room] = keys}
  visits = Array.new(rooms.size, false)
  visit(0, rooms, visits)
  visits.all?{|visit| visit}
end

def visit(room, rooms, visits)
  return if visits[room]
  visits[room] = true
  rooms[room].each{|key| visit(key, rooms, visits)}
  visits
end

rooms = [[1],[2],[3],[]]
p can_visit_all_rooms(rooms)


rooms = [[1,3],[3,0,1],[2],[0]]
p can_visit_all_rooms(rooms)
