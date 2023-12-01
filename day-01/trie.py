class Node:
  def __init__(self) -> None:
    self.children = dict()
    self.word = ""


class Trie:
  def __init__(self) -> None:
    self.root = Node()

  def insert(self, word: str) -> None:
    node = self.root
    
    for c in word:
      if c not in node.children:
        node.children[c] = Node()
      node = node.children[c]
    node.word = word

  def find(self, idx: int, string: str) -> str:
    
    cur_node = self.root
    for i in range(idx, len(string)):
      if string[i] not in cur_node.children:
        return ""
      cur_node = cur_node.children[string[i]]
      
      if (cur_node.word) > 0:
        return cur_node.word
    
    return ""


trie = Trie()
