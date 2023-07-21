
import sys
import grafo
import biblioteca
import funciones as f
from collections import deque

def main():
    argumentos = sys.argv
    argumentos = argumentos[1:]
    if len(argumentos) != 2:
        raise Exception("Parametros invalidos")
    aeropuertos = argumentos[0]
    vuelos = argumentos[1]
    pila = deque()

    grafo_dinero = grafo.Grafo()
    grafo_tiempos = grafo.Grafo()
    grafo_frecuencias = grafo.Grafo()
    ciudades = {} #{ciudad: [aeropuerto1, aeropuerto2], ciudad2:[aeropuerto3, aeropuerto4]}
    coordenadas = {}

    f.ingresar_aeropuertos(aeropuertos, grafo_dinero, grafo_tiempos, grafo_frecuencias, ciudades, coordenadas)
    f.ingresar_vuelos(vuelos, grafo_dinero, grafo_tiempos, grafo_frecuencias)

    for linea in sys.stdin:
        comando, parametros = linea.split(" ", 1)
        parametros = parametros.rstrip().split(",")

        if comando == "camino_mas":
            if len(parametros) != 3:
                print("Error de parametros")
                continue

            criterio, origen, destino = parametros[0], parametros[1], parametros[2]

            if criterio == "barato":
                g = grafo_dinero
            elif criterio == "rapido":
                g = grafo_tiempos
            else:
                print("Error de parametros")
                continue
            
            if origen not in ciudades or destino not in ciudades:
                print("Error de parametros")
                continue
            
            res = f.camino_minimo(origen, destino, ciudades, g)
            pila.append(res)
            print(" -> ".join(res))

        elif comando == "camino_escalas":
            if len(parametros) != 2:
                print("Error de parametros")
                continue

            origen, destino = parametros[0], parametros[1]
            g = grafo_dinero

            if origen not in ciudades or destino not in ciudades:
                print("Error de parametros")
                continue

            res = f.camino_minimo_escalas(origen, destino, ciudades, g)
            pila.append(res)
            print(" -> ".join(res))

        elif comando == "centralidad":
            if len(parametros) != 1:
                print("Error de parametros")
                continue

            n = parametros[0]
            if not n.isdigit():
                print("Error de parametros")
                continue
            
            n = int(n)
            g = grafo_frecuencias

            res = f.obtener_centralidad(g, n)
            print(", ".join(res))

        elif comando == "itinerario":
            if len(parametros) != 1:
                print("Error de parametros")

            ruta = parametros[0]
            grafo_ciudades = f.crear_itinerario(ruta)
            orden = biblioteca.orden_topologico_dfs(grafo_ciudades)
            print(", ".join(orden))

            for i in range(len(orden)-1):
                origen, destino = orden[i], orden[i+1]
                res = f.camino_minimo_escalas(origen, destino, ciudades, grafo_frecuencias)
                print(" -> ".join(res))
        
        elif comando == "nueva_aerolinea":
            if len(parametros) != 1:
                print("Error de parametros")
                continue
            ruta = parametros[0]
            arbol = biblioteca.mst_prim(grafo_dinero)
            f.crear_rutas(ruta, arbol, grafo_tiempos, grafo_frecuencias)
            print("OK")
        
        elif comando == "exportar_kml":
            if len(parametros) != 1:
                print("Error de parametros")
                continue
            ruta = parametros[0]
            if len(pila) == 0:
                print("No se ha ejecutado un comando anteriormente")
                continue
            camino = pila.pop()
            f.crear_kml(ruta, camino, coordenadas)
            print("OK")
        
main()