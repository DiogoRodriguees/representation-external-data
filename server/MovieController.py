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
        
    def findByCategories(self, request):
        try:
            print("[Movie Controller] Executing method findByCategories()")
            response = self.movieService.findByCategories(request.filters.values)
            return Response(status=200, message="Successfuly on find movie by categories", movies=response)
        except Exception as e:
            print(f"[Error] Failed on find movie by categories{e}")
            return Response(status=400, message="Failed on find movie by categories: " + str(e))
        
    def findByAtor(self, request):
        try:
            print("[Movie Controller] Executing method findByAtor()")
            response = self.movieService.findByAtor(request.filters.values)
            return Response(status=200, message="Successfuly on find movie by ator", movies=response)
        except Exception as e:
            print(f"[Error] Failed on find movie by actors{e}")
            return Response(status=400, message="Failed on find movie by actor: " + str(e))
