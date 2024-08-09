from Movies_pb2 import Response

class MovieController:
    def __init__(self, movieService):
        self.movieService = movieService

    def create(self, request):
        try:
            print("[Movie Controller] Executing method create()")
            movie = self.movieService.create(request.movie)
            return Response(status=200, message="Successfuly on create movie", movie=movie)
        except Exception as e:
            print(f"[Error] Failed on create movie: {e}")
            return Response(status=400, message="Failed on create movie: " + str(e))
        
    def update(self, request):
        try:
            print("[Movie Controller] Executing method update()")
            movie = self.movieService.update(request.movie)
            return Response(status=200, message="Successfully on update movie")
        except Exception as e:
            print(f"[Error] Failed on update movie: {e}")
            return Response(status=400, message="Failed on update movie: " + str(e))
        
        
    def findByCategories(self, request):
        try:
            print("[Movie Controller] Executing method findByCategories()")
            response = self.movieService.findByCategories(request.filters.values)
            return Response(status=200, message="Successfully on find movie by categories", movies=response)
        except Exception as e:
            print(f"[Error] Failed on find movie by categories{e}")
            return Response(status=400, message="Failed on find movie by categories: " + str(e))
        
    def findByAtor(self, request):
        try:
            print("[Movie Controller] Executing method findByAtor()")
            response = self.movieService.findByAtor(request.filters.values)
            return Response(status=200, message="Successfully on find movie by ator", movies=response)
        except Exception as e:
            print(f"[Error] Failed on find movie by actors{e}")
            return Response(status=400, message="Failed on find movie by actor: " + str(e))
    
    def delete(self, request):
        try:
            print("[Movie Controller] Executing method delete()")
            self.movieService.delete(request.movie)
            return Response(status=200, message="Successfully on delete movie with id: ")
        except Exception as e:
            print(f"[Error] Failed on delete movie: {e}")
            return Response(status=400, message="Failed on delete movie" + str(e))
