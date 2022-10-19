# @param {String} path
# @return {String}
def simplify_path(path)
  stack = []
  path.split('/') do |file|
    next if file.empty? || file == '.'
    if file == '..'
      stack.pop
    else
      stack.push file
    end
  end
  '/' << stack.join('/')
end

path = "/home/"
p simplify_path(path)
path = "/../"
p simplify_path(path)
path = "/home//foo/"
p simplify_path(path)
path = "/a/./b/../../c/"
p simplify_path(path)
