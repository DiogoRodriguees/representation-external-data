import socket
import struct

import movie
from movie.Movies_pb2 import Request
from database.Database import Database
from movie.MovieService import MovieService
from movie.MovieController import MovieController


class Server:
    def __init__(self):
        self.database = Database()
        self.movieService = MovieService(self.database)
        self.movieController = MovieController(self.movieService)
        self.server_socket = None

    def middleware(self, request):
        print("Using middleware, method: ", request.method)
        if request.method == "CREATE":
            return self.movieController.create(request)
        if request.method == "FIND_BY_ATOR":
            return self.movieController.find_by_actor(request)
        if request.method == "FIND_BY_CATEGORIA":
            return self.movieController.find_by_categories(request)
        if request.method == "DELETE":
            return self.movieController.delete(request)
        if request.method == "UPDATE":
            return self.movieController.update(request)

    def handle_client(self, client_socket):
        try:
            # Read data size
            size_data = client_socket.recv(4)
            size = struct.unpack("!I", size_data)[0]

            # Read payload
            data = client_socket.recv(size)

            # Unmarshalling Request
            request = Request()
            request.ParseFromString(data)

            # Process request
            response = self.middleware(request)

            # Marshalling response
            response_data = response.SerializeToString()

            # Send response size followd by response
            response_size = struct.pack("!I", len(response_data))
            client_socket.sendall(response_size)
            client_socket.sendall(response_data)
        except Exception as e:
            print(f"Erro ao processar requisição: {e}")
            raise e

    def run(self):
        server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server.bind(("localhost", 8080))
        server.listen(1)
        self.server_socket = server

        print("Server running...")

        try:
            client_socket, addr = server.accept()
            print(f"Conexão estabelecida com {addr}")
            while True:
                try:
                    self.handle_client(client_socket)
                except Exception as e:
                    raise e

        except Exception as e:
            print("Error: ", e)

        finally:
            client_socket.close()
            if self.server_socket:
                self.server_socket.close()


server = Server()
server.run()
