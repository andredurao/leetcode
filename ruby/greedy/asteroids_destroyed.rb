# @param {Integer} mass
# @param {Integer[]} asteroids
# @return {Boolean}
def asteroids_destroyed(mass, asteroids)
  asteroids.sort.each do |asteroid|
    if mass >= asteroid
      mass += asteroid
    else
      return false
    end
  end
  true
end
