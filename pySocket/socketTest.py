import socket
import time
HOST = '192.168.1.69'
PORT = 80


with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as s:
    s.connect((HOST, PORT))
    total = 0 
    n=0
    while n<10000:
        start = time.perf_counter()
        s.sendall(b'hello')
        data = s.recv(1024)
        total+= time.perf_counter()-start
        n+=1
    print("Avg time ->"+str(total/n))