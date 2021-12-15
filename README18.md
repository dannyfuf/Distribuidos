# Distribuidos

En este nodo solo se ejecuta un Fulcrum server y el cliente de Leia.

Al ingresar a la maquina:
    - Ingresar a `LAB3`
    - Ejecutar `go mod tidy`

Se asume un terminal conectado a la máquina 18:

1. Ejecutar Fulcrum Server
    - Ejecutar `make server name=fulcrum`

2. Ejecutar Leia
   - Ejecutar `make leia`

3. Consideraciones:
    - Los string ingresados por consola  no pueden contener espacios, ni comas (" ", ",")
    - Al ingresar al servidor fulcrum, se pedirá el numero de la maquina desde la cual se conectó. En esta máquina virtual, se debe ingresar `18`

# Integrantes
    - Hans Fluhmann 201773565-1
    - Danny Fuentes 201773559-7
    - Rodrigo Orellana 201773522-8