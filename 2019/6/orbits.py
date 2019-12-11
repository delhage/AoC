import os
def commit(parent, name):
    if parent:
        os.system("git checkout " + parent)
    os.system("git commit -m " + name + " --allow-empty")
    os.system("git tag " + name)

tree = {}
for line in open("./input6.txt", "r").read().splitlines():
    center,orbiter = line.split(")")
    tree[orbiter] = center

queue = [(None, "COM")]
while queue:
    parent,el = queue.pop()
    commit(parent, el)
    queue.extend((el, t) for t in tree if tree[t] == el)

