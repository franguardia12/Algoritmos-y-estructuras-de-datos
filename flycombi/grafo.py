class Grafo:
    def __init__(self, es_dirigido = False, vertices = []):
        self.dirigido = es_dirigido
        self.vertices = {}
        if len(vertices) != 0:
            for vertice in vertices:
                self.vertices[vertice] = {}
    
    def agregar_vertice(self, v):
        if v in self.vertices:
            raise ValueError(f"Ya hay un vertice {v} en el grafo")
        self.vertices[v] = {}
    
    def borrar_vertice(self, v):
        if v not in self.vertices:
            raise ValueError(f"No hay un vertice {v} en el grafo")
        self.vertices.pop(v)
        for adyacentes in self.vertices.values():
            for vertice in adyacentes.keys():
                if vertice == v:
                    adyacentes.pop(v)
    
    def agregar_arista(self, v, w, peso = 1):
        if v not in self.vertices:     
            raise ValueError(f"No hay un vertice {v} en el grafo")
        elif w not in self.vertices:
            raise ValueError(f"No hay un vertice {w} en el grafo")
        elif w in self.vertices[v]:
            raise ValueError(f"El vertice {v} ya tiene como adyacente al vertice {w}")
        
        if not self.dirigido:
            self.vertices[w][v] = peso
        self.vertices[v][w] = peso
    
    def estan_unidos(self, v, w):
        if not self.dirigido:
            if w not in self.vertices[v]:
                return False
        if v not in self.vertices or w not in self.vertices or not v in self.vertices[w]:
            return False
        return True
    
    def peso_arista(self, v, w):
        if not self.estan_unidos(v,w):
            raise ValueError(f"El vertice {v} no tiene como adyacente el vertice {w}")
        return self.vertices[v][w]

    def obtener_vertices(self):
        res = []
        for vertice in self.vertices.keys():
            res.append(vertice)
        return res
    
    def vertice_aleatorio(self):
        vertices = self.obtener_vertices()
        if len(vertices) == 0:
            return None
        return vertices[0]

    def adyacentes(self, v):
        res = []
        for adyacente in self.vertices[v].keys():
            res.append(adyacente)
        return res
    