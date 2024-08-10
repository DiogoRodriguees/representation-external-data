from movie.Movies_pb2 import Response
from movie.Movie import MovieDTO


class MovieController:
    def __init__(self, movieService):
        self.movieService = movieService

    def create(self, request):
        try:
            print("[Movie Controller] Executing method create()")
            movie = self.movieService.create(request.movie)
            msg = "Successfuly on create movie"
            return Response(status=200, message=msg, movie=movie)
        except Exception as e:
            print(f"[Error] Failed on create movie: {e}")
            return
            return Response(status=400, message="Failed on create movie: " + str(e))

    def update(self, request):
        try:
            print("[Movie Controller] Executing method update()")
            movie = self.movieService.update(request.movie)
            return Response(status=200, message="Successfully on update movie")
        except Exception as e:
            print(f"[Error] Failed on update movie: {e}")
            return Response(status=400, message="Failed on update movie: " + str(e))

    def find_by_categories(self, request):
        try:
            print("[Movie Controller] Executing method findByCategories()")
            movies = self.movieService.find_by_categories(request.filters.values)
            mes = "Successfully on find by categories"
            return Response(status=200, message=msg, movies=movies)
        except Exception as e:
            print(f"[Error] Failed on find movie by categories: {e}")
            return Response(status=400, message="Failed on find by genres: " + str(e))

    def find_by_actor(self, request):
        try:
            print("[Movie Controller] Executing method findByAtor()")
            movies = self.movieService.finc_by_actor(request.filters.values)
            msg = "Successfully on find movie by ator"
            return Response(status=200, message=msg, movies=movies)
        except Exception as e:
            print(f"[Error] Failed on find movie by actors: {e}")
            return Response(status=400, message="Failed on find by actor: " + str(e))

    def delete(self, request):
        try:
            print("[Movie Controller] Executing method delete()")
            self.movieService.delete(request.movie)
            id = str(request.movie.id)
            return Response(status=200, message="Successfully on delete with id: " + id)
        except Exception as e:
            print(f"[Error] Failed on delete movie: {e}")
            return Response(status=400, message="Failed on delete movie" + str(e))
